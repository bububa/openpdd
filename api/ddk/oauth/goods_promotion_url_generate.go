package oauth

import (
	"github.com/bububa/openpdd/api/ddk"
	"github.com/bububa/openpdd/core"
)

// GoodsPromotionUrlGenerateRequest 多多进宝推广链接生成 API Request
type GoodsPromotionUrlGenerateRequest struct {
	ddk.GoodsPromotionUrlGenerateRequest
	// ForceDuoID 是否使用多多客专属推广计划
	ForceDuoID bool `json:"force_duo_id,omitempty"`
}

// GetType implement Request interface
func (r GoodsPromotionUrlGenerateRequest) GetType() string {
	return "pdd.ddk.oauth.goods.promotion.url.generate"
}

// GoodsPromotionUrlGenerate 多多进宝推广链接生成
func GoodsPromotionUrlGenerate(clt *core.SDKClient, req *GoodsPromotionUrlGenerateRequest, accessToken string) ([]ddk.PromURL, error) {
	var resp ddk.GoodsPromotionUrlGenerateResponse
	if err := clt.Do(req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Response.List, nil
}
