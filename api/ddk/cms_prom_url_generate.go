package ddk

import (
	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// CmsPromUrlGenerateRequest 生成商城-频道推广链接 API Request
type CmsPromUrlGenerateRequest struct {
	// ChannelType 0, "1.9包邮"；1, "今日爆款"； 2, "品牌清仓"； 3, "默认商城"；4,"PC端专属商城"；非必填 ,默认是3,
	ChannelType *int `json:"channel_type,omitempty"`
	// CustomParameters 自定义参数，为链接打上自定义标签。自定义参数最长限制64个字节。
	CustomParameters interface{} `json:"custom_parameters,omitempty"`
	// GenerateMobile 是否生成手机跳转链接。true-是，false-否，默认false
	GenerateMobile bool `json:"generate_mobile,omitempty"`
	// GenerateSchemeURL 是否返回 schema URL
	GenerateSchemeURL bool `json:"generate_scheme_url,omitempty"`
	// GenerateShortURL 是否生成短链接，true-是，false-否
	GenerateShortURL bool `json:"generate_short_url,omitempty"`
	// GenerateWeApp 是否生成拼多多福利券微信小程序推广信息
	GenerateWeApp bool `json:"generate_we_app,omitempty"`
	// Keyword 搜索关键词
	Keyword string `json:"keyword,omitempty"`
	//单人团多人团标志。true-多人团，false-单人团 默认false
	MultiGroup bool `json:"multi_group,omitempty"`
	//推广位列表，例如：["60005_612"]
	PidList string `json:"p_id_list,omitempty"`
}

// GetType implement Request interface
func (r CmsPromUrlGenerateRequest) GetType() string {
	return "pdd.ddk.cms.prom.url.generate"
}

// CmsPromUrlGenerateResponse 生成商城-频道推广链接 API Response
type CmsPromUrlGenerateResponse struct {
	model.CommonResponse
	// Response 商城推广链接返回对象
	Response struct {
		// Total total
		Total int `json:"total,omitempty"`
		// URLList 链接列表
		URLList []PromURL `json:"url_list,omitempty"`
	} `json:"cms_promotion_url_generate_response" xml:"cms_promotion_url_generate_response"`
}

// CmsPromUrlGenerate 生成商城-频道推广链接
func CmsPromUrlGenerate(clt *core.SDKClient, req *CmsPromUrlGenerateRequest) (int, []PromURL, error) {
	var resp CmsPromUrlGenerateResponse
	if err := clt.Do(req, &resp, ""); err != nil {
		return 0, nil, err
	}
	return resp.Response.Total, resp.Response.URLList, nil
}
