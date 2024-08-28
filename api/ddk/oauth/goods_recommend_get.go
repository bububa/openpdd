package oauth

import (
	"context"

	"github.com/bububa/openpdd/api/ddk"
	"github.com/bububa/openpdd/core"
)

// GoodsRecommendGetRequest 多多进宝商品推荐API API Request
type GoodsRecommendGetRequest struct {
	ddk.GoodsRecommendGetRequest
	// ForceAuthDuoID 是否使用工具商专属推广计划，默认为false
	ForceAuthDuoID bool `json:"force_auth_duo_id,omitempty"`
}

// GetType implement Request interface
func (r GoodsRecommendGetRequest) GetType() string {
	return "pdd.ddk.oauth.goods.recommend.get"
}

// GoodsRecommendGet 多多进宝商品推荐API
func GoodsRecommendGet(ctx context.Context, clt *core.SDKClient, req *GoodsRecommendGetRequest, accessToken string) (*ddk.GoodsRecommendGetResult, error) {
	var resp ddk.GoodsRecommendGetResponse
	if err := clt.Do(ctx, req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Response, nil
}
