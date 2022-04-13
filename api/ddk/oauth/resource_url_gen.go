package oauth

import (
	"github.com/bububa/openpdd/api/ddk"
	"github.com/bububa/openpdd/core"
)

// ResourceUrlGenRequest 生成多多进宝频道推广 API Request
type ResourceUrlGenRequest struct {
	// CustomParameters 自定义参数，为链接打上自定义标签；自定义参数最长限制64个字节；格式为： {"uid":"11111","sid":"22222"} ，其中 uid 用户唯一标识，可自行加密后传入，每个用户仅且对应一个标识，必填； sid 上下文信息标识，例如sessionId等，非必填。该json字符串中也可以加入其他自定义的key
	CustomParameters interface{} `json:"custom_parameters,omitempty"`
	// GenerateSchemeURL 是否返回 schema URL
	GenerateSchemeURL bool `json:"generate_scheme_url,omitempty"`
	// GenerateWeApp 是否生成拼多多福利券微信小程序推广信息
	GenerateWeApp bool `json:"generate_we_app,omitempty"`
	// Pid 推广位
	Pid string `json:"pid,omitempty"`
	// ResourceType 频道来源：4-限时秒杀,39997-充值中心, 39998-活动转链，39996-百亿补贴，39999-电器城，40000-领券中心，50005-火车票
	ResourceType int `json:"resource_type,omitempty"`
	// URL 原链接
	URL string `json:"url,omitempty"`
}

// GetType implement Request interface
func (r ResourceUrlGenRequest) GetType() string {
	return "pdd.ddk.oauth.resource.url.gen"
}

// ResourceUrlGen 生成多多进宝频道推广
func ResourceUrlGen(clt *core.SDKClient, req *ResourceUrlGenRequest, accessToken string) (*ddk.PromURL, error) {
	var resp ddk.ResourceUrlGenResponse
	if err := clt.Do(req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Response, nil
}
