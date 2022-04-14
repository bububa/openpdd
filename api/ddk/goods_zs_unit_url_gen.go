package ddk

import (
	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// GoodsZsUnitUrlGenRequest 多多进宝转链接口 API Request
type GoodsZsUnitUrlGenRequest struct {
	// CustomParameters 自定义参数，为链接打上自定义标签；自定义参数最长限制64个字节；格式为： {"uid":"11111","sid":"22222"} ，其中 uid 用户唯一标识，可自行加密后传入，每个用户仅且对应一个标识，必填； sid 上下文信息标识，例如sessionId等，非必填。该json字符串中也可以加入其他自定义的key。（如果使用GET请求，请使用URLEncode处理参数）
	CustomParameters interface{} `json:"custom_parameters,omitempty"`
	// Pid 渠道id
	Pid string `json:"pid"`
	// SourceURL 需转链的链接，支持拼多多商品链接、进宝长链/短链（即为pdd.ddk.goods.promotion.url.generate接口生成的长短链）
	SourceURL string `json:"source_url"`
}

// GetType implement Request interface
func (r GoodsZsUnitUrlGenRequest) GetType() string {
	return "pdd.ddk.goods.zs.unit.url.gen"
}

// GoodsZsUnitUrlGenResponse 多多进宝转链接口 API REesponse
type GoodsZsUnitUrlGenResponse struct {
	model.CommonResponse
	// Response response
	Response *PromURL `json:"goods_zs_unit_generate_response" xml:"goods_zs_unit_generate_response"`
}

// GoodsZsUnitUrlGen 多多进宝转链接口
func GoodsZsUnitUrlGen(clt *core.SDKClient, req *GoodsZsUnitUrlGenRequest) (*PromURL, error) {
	var resp GoodsZsUnitUrlGenResponse
	if err := clt.Do(req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Response, nil
}
