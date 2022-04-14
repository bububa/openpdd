package oauth

import (
	"github.com/bububa/openpdd/api/ddk"
	"github.com/bububa/openpdd/core"
)

// GoodsDetailRequest 多多进宝商品详情查询 API Request
type GoodsDetailRequest struct {
	ddk.GoodsDetailRequest
}

// GetType implement Request interface
func (r GoodsDetailRequest) GetType() string {
	return "pdd.ddk.oauth.goods.detail"
}

// GoodsDetail 多多进宝商品详情查询
func GoodsDetail(clt *core.SDKClient, req *GoodsDetailRequest, accessToken string) ([]ddk.Goods, error) {
	var resp ddk.GoodsDetailResponse
	if err := clt.Do(req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Response.List, nil
}
