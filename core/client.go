package core

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/bububa/openpdd/core/internal/debug"
	"github.com/bububa/openpdd/model"
	"github.com/bububa/openpdd/util/query"
)

// SDKClient sdk client
type SDKClient struct {
	clientID string
	secret   string
	dataType model.RequestDataType
	gw       string
	debug    bool
}

// NewSDKClient 创建SDKClient
func NewSDKClient(clientID string, secret string) *SDKClient {
	return &SDKClient{
		clientID: clientID,
		secret:   secret,
		gw:       GATEWAY,
		dataType: model.RequestDataType_JSON,
	}
}

// SetDebug 设置debug模式
func (c *SDKClient) SetDebug(debug bool) {
	c.debug = debug
}

// SetGateway 设置gateway
func (c *SDKClient) SetGateway(gw string) {
	c.gw = gw
}

// SetDataType 设置返回数据格式
func (c *SDKClient) SetDataType(t model.RequestDataType) {
	c.dataType = t
}

// ClientID returns client client_id
func (c *SDKClient) ClientID() string {
	return c.clientID
}

func (c *SDKClient) Do(req model.Request, resp model.Response, accessToken string) error {
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
	return c.get(values, resp)
}

func (c *SDKClient) post(req url.Values, resp model.Response) error {
	var builder strings.Builder
	builder.WriteString(c.gw)
	if req != nil {
		builder.WriteString("?")
		builder.WriteString(req.Encode())
	}
	reqUrl := builder.String()
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequest("POST", c.gw, strings.NewReader(req.Encode()))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	return c.fetch(httpReq, resp)
}

func (c *SDKClient) get(req url.Values, resp model.Response) error {
	var builder strings.Builder
	builder.WriteString(c.gw)
	if req != nil {
		builder.WriteString("?")
		builder.WriteString(req.Encode())
	}
	reqUrl := builder.String()
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return err
	}
	return c.fetch(httpReq, resp)
}

// fetch execute http request
func (c *SDKClient) fetch(httpReq *http.Request, resp model.Response) error {
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	if resp == nil {
		resp = &model.CommonResponse{}
	}
	err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	if resp.IsError() {
		return resp.Error()
	}
	return nil
}

func (c *SDKClient) sign(values url.Values) string {
	params := make([]string, 0, len(values))
	for k := range values {
		var builder strings.Builder
		builder.WriteString(k)
		builder.WriteString(values.Get(k))
		params = append(params, builder.String())
	}
	sort.Strings(params)
	var builder strings.Builder
	builder.WriteString(c.secret)
	for _, v := range params {
		builder.WriteString(v)
	}
	builder.WriteString(c.secret)
	rawSign := builder.String()
	return strings.ToUpper(md5String(rawSign))
}

func md5String(src string) string {
	h := md5.New()
	io.WriteString(h, src)
	return hex.EncodeToString(h.Sum(nil))
}
