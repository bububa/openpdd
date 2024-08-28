package core

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/bububa/openpdd/core/internal/debug"
	"github.com/bububa/openpdd/model"
	"github.com/bububa/openpdd/util"
	"github.com/bububa/openpdd/util/query"
)

var (
	onceInit   sync.Once
	httpClient *http.Client
)

func defaultHttpClient() *http.Client {
	onceInit.Do(func() {
		transport := http.DefaultTransport.(*http.Transport).Clone()
		transport.MaxIdleConns = 100
		transport.MaxConnsPerHost = 100
		transport.MaxIdleConnsPerHost = 100
		httpClient = &http.Client{
			Transport: transport,
			Timeout:   time.Second * 60,
		}
	})
	return httpClient
}

// SDKClient sdk client
type SDKClient struct {
	client   *http.Client
	tracer   *Otel
	clientID string
	secret   string
	dataType model.RequestDataType
	gw       string
	uploadGw string
	debug    bool
}

// NewSDKClient 创建SDKClient
func NewSDKClient(clientID string, secret string) *SDKClient {
	return &SDKClient{
		clientID: clientID,
		secret:   secret,
		gw:       GATEWAY,
		uploadGw: UPLOAD_GATEWAY,
		dataType: model.RequestDataType_JSON,
		client:   defaultHttpClient(),
	}
}

// SetDebug 设置debug模式
func (c *SDKClient) SetDebug(debug bool) {
	c.debug = debug
}

// SetHttpClient 设置http.Client
func (c *SDKClient) SetHttpClient(client *http.Client) {
	c.client = client
}

// SetGateway 设置gateway
func (c *SDKClient) SetGateway(gw string) {
	c.gw = gw
}

func (c *SDKClient) SetUploadGateway(gw string) {
	c.uploadGw = gw
}

// SetDataType 设置返回数据格式
func (c *SDKClient) SetDataType(t model.RequestDataType) {
	c.dataType = t
}

func (c *SDKClient) WithTracer(namespace string) {
	c.tracer = NewOtel(namespace, c.ClientID())
}

// ClientID returns client client_id
func (c *SDKClient) ClientID() string {
	return c.clientID
}

func (c *SDKClient) Do(ctx context.Context, req model.Request, resp model.Response, accessToken string) error {
	values, err := query.Values(req)
	if err != nil {
		return err
	}
	values.Set("type", req.GetType())
	values.Set("client_id", c.clientID)
	values.Set("timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	values.Set("data_type", string(c.dataType))
	if accessToken != "" {
		values.Set("access_token", accessToken)
	}
	values.Set("sign", c.sign(values))
	return c.get(ctx, req.GetType(), values, resp)
}

func (c *SDKClient) Upload(ctx context.Context, req model.UploadRequest, resp model.Response, accessToken string) error {
	fields := req.Encode()
	fields = append(fields, []model.UploadField{
		{
			Key:   "type",
			Value: req.GetType(),
		}, {
			Key:   "client_id",
			Value: c.clientID,
		}, {
			Key:   "timestamp",
			Value: strconv.FormatInt(time.Now().Unix(), 10),
		}, {
			Key:   "date_type",
			Value: string(c.dataType),
		},
	}...)
	if accessToken != "" {
		fields = append(fields, model.UploadField{
			Key:   "access_token",
			Value: accessToken,
		})
	}
	fields = append(fields, model.UploadField{
		Key:   "sign",
		Value: c.signUploadFields(fields),
	})
	return c.upload(ctx, req.GetType(), fields, resp)
}

func (c *SDKClient) post(ctx context.Context, methodName string, req url.Values, resp model.Response) error {
	payload := req.Encode()
	var builder strings.Builder
	builder.WriteString(c.gw)
	if req != nil {
		builder.WriteString("?")
		builder.WriteString(payload)
	}
	reqUrl := builder.String()
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.gw, strings.NewReader(payload))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	return c.WithSpan(ctx, methodName, httpReq, resp, []byte(payload), c.fetch)
}

func (c *SDKClient) get(ctx context.Context, methodName string, req url.Values, resp model.Response) error {
	var builder strings.Builder
	builder.WriteString(c.gw)
	if req != nil {
		builder.WriteString("?")
		builder.WriteString(req.Encode())
	}
	reqUrl := builder.String()
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl, nil)
	if err != nil {
		return err
	}
	return c.WithSpan(ctx, methodName, httpReq, resp, nil, c.fetch)
}

func (c *SDKClient) upload(ctx context.Context, methodName string, req []model.UploadField, resp model.Response) error {
	var buf bytes.Buffer
	var builder strings.Builder
	mw := multipart.NewWriter(&buf)
	mp := make(map[string]string, len(req))
	for _, f := range req {
		var (
			fw  io.Writer
			r   io.Reader
			err error
		)
		if f.Reader != nil {
			if fw, err = mw.CreateFormFile(f.Key, f.Value); err != nil {
				return err
			}
			r = f.Reader
			builder.WriteString("@")
			if f.Value == "" {
				f.Value = "/tmp"
			}
			builder.WriteString(f.Value)
			mp[f.Key] = builder.String()
			builder.Reset()
		} else {
			if fw, err = mw.CreateFormField(f.Key); err != nil {
				return err
			}
			r = strings.NewReader(f.Value)
			mp[f.Key] = f.Value
		}
		if _, err = io.Copy(fw, r); err != nil {
			return err
		}
	}
	mw.Close()
	debug.PrintPostMultipartRequest(c.uploadGw, mp, c.debug)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uploadGw, &buf)
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", mw.FormDataContentType())
	bs, _ := json.Marshal(mp)
	return c.WithSpan(ctx, methodName, httpReq, resp, bs, c.fetch)
}

// fetch execute http request
func (c *SDKClient) fetch(httpReq *http.Request, resp model.Response) (*http.Response, error) {
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return httpResp, err
	}
	defer httpResp.Body.Close()
	if resp == nil {
		resp = &model.CommonResponse{}
	}
	err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return httpResp, err
	}
	if resp.IsError() {
		return httpResp, resp.Error()
	}
	return httpResp, nil
}

func (c *SDKClient) sign(values url.Values) string {
	params := make([]string, 0, len(values))
	var builder strings.Builder
	for k := range values {
		builder.WriteString(k)
		builder.WriteString(values.Get(k))
		params = append(params, builder.String())
		builder.Reset()
	}
	sort.Strings(params)
	builder.WriteString(c.secret)
	for _, v := range params {
		builder.WriteString(v)
	}
	builder.WriteString(c.secret)
	rawSign := builder.String()
	return strings.ToUpper(util.Md5String(rawSign))
}

func (c *SDKClient) signUploadFields(fields []model.UploadField) string {
	params := make([]string, 0, len(fields))
	var builder strings.Builder
	for _, f := range fields {
		if f.Reader != nil {
			continue
		}
		builder.WriteString(f.Key)
		builder.WriteString(f.Value)
		params = append(params, builder.String())
		builder.Reset()
	}
	sort.Strings(params)
	builder.WriteString(c.secret)
	for _, v := range params {
		builder.WriteString(v)
	}
	builder.WriteString(c.secret)
	rawSign := builder.String()
	return strings.ToUpper(util.Md5String(rawSign))
}

func (c *SDKClient) WSSUrl() string {
	ts := strconv.FormatInt(time.Now().UnixMilli(), 10)
	var builder strings.Builder
	builder.WriteString(WSSEndPoint)
	builder.WriteString("/")
	builder.WriteString(c.clientID)
	builder.WriteString("/")
	builder.WriteString(ts)
	builder.WriteString("/")
	builder.WriteString(c.signWSS(ts))
	return builder.String()
}

func (c *SDKClient) signWSS(ts string) string {
	var builder strings.Builder
	builder.WriteString(c.clientID)
	builder.WriteString(ts)
	builder.WriteString(c.secret)
	return base64.StdEncoding.EncodeToString([]byte(util.Md5String(builder.String())))
}

func (c *SDKClient) WithSpan(ctx context.Context, methodName string, req *http.Request, resp model.Response, payload []byte, fn func(*http.Request, model.Response) (*http.Response, error)) error {
	if c.tracer == nil {
		_, err := fn(req, resp)
		return err
	}
	return c.tracer.WithSpan(ctx, methodName, req, resp, payload, fn)
}
