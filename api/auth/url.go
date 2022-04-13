package auth

import (
	"net/url"
	"strings"

	"github.com/bububa/openpdd/core"
)

// URLRequest 组装授权页URL API Request
type URLRequest struct {
	// AuthType 授权类型
	AuthType AuthType `json:"auth_type"`
	// RedirectURI 已创建应用-应用详情中查看回调地址字段值，在用户完成授权后，code值将回调至该地址中。注意：redirect_uri需进行url编码（encodeURIComponent）
	RedirectURI string `json:"redirect_uri,omitempty"`
	// State 自定义参数，授权完成后，原值将回调至redirect_uri中
	State string `json:"state,omitempty"`
	// View 可选web或h5，默认为web，H5移动端网页授权必填
	View string `json:"view,omitempty"`
}

// URL 组装授权页URL
func URL(clt *core.SDKClient, req *URLRequest) string {
	var builder strings.Builder
	switch req.AuthType {
	case AuthType_SHOP_WEB:
		builder.WriteString(SHOP_AUTH_WEB_URL)
	case AuthType_SHOP_H5:
		builder.WriteString(SHOP_AUTH_H5_URL)
	case AuthType_DDK:
		builder.WriteString(DDK_AUTH_URL)
	case AuthType_KTT:
		builder.WriteString(KTT_AUTH_URL)
	case AuthType_LOGISTIC:
		builder.WriteString(LOGISTIC_AUTH_URL)
	}
	builder.WriteString("?")
	values := &url.Values{}
	values.Set("client_id", clt.ClientID())
	values.Set("response_type", "code")
	values.Set("redirect_uri", req.RedirectURI)
	if req.State != "" {
		values.Set("state", req.State)
	}
	if req.AuthType == AuthType_SHOP_H5 {
		values.Set("view", H5)
	} else if req.View != "" {
		values.Set("view", req.View)
	}
	builder.WriteString(values.Encode())
	return builder.String()
}
