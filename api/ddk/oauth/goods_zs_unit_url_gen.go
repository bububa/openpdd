package oauth

import (
	"github.com/bububa/openpdd/api/ddk"
	"github.com/bububa/openpdd/core"
)

// GoodsZsUnitUrlGenRequest 多多进宝转链接口 API Request
type GoodsZsUnitUrlGenRequest struct {
	ddk.GoodsZsUnitUrlGenRequest
	// GenerateSchemeURL 是否返回 schema URL
	GenerateSchemeURL bool `json:"generate_scheme_url,omitempty"`
}

// GetType implement Request interface
func (r GoodsZsUnitUrlGenRequest) GetType() string {
	return "pdd.ddk.oauth.goods.zs.unit.url.gen"
}

// GoodsZsUnitUrlGen 多多进宝转链接口
func GoodsZsUnitUrlGen(clt *core.SDKClient, req *GoodsZsUnitUrlGenRequest, accessToken string) (*ddk.PromURL, error) {
	var resp ddk.GoodsZsUnitUrlGenResponse
	if err := clt.Do(req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Response, nil
}
