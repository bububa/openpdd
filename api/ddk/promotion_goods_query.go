package ddk

import (
	"context"

	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// PromotionGoodsQueryRequest 多多进宝信息流投放商品报备进度查询 API Request
type PromotionGoodsQueryRequest struct {
	// PageNumber 分页查询页数
	PageNumber int `json:"page_number,omitempty"`
	// PageSize 分页查询页大小
	PageSize int `json:"page_size,omitempty"`
	// GoodsID 商品id
	GoodsID uint64 `json:"goods_id,omitempty"`
	// MallID 店铺id
	MallID uint64 `json:"mall_id,omitempty"`
	// StatusList 查询状态列表
	StatusList []int `json:"status_list,omitempty"`
	// UpdateEndTime 最后更新结束时间（最长支持30天）
	UpdateEndTime int64 `json:"update_end_time,omitempty"`
	// UpdateStartTime 最后更新开始时间
	UpdateStartTime int64 `json:"update_start_time,omitempty"`
}

// GetType implement Request interface
func (r PromotionGoodsQueryRequest) GetType() string {
	return "pdd.ddk.promotion.goods.query"
}

// PromotionGoodsQueryResponse 多多进宝信息流投放商品报备进度查询 API Response
type PromotionGoodsQueryResponse struct {
	model.CommonResponse
	Response struct {
		// List 报备记录列表
		List []PromotionGoodsApplication `json:"application_list,omitempty"`
		// Total 报备记录总数
		Total int `json:"total,omitempty"`
	} `json:"response"`
}

// PromotionGoodsApplication 多多进宝信息流投放商品报备进度
type PromotionGoodsApplication struct {
	// Comment 审核信息
	Comment string `json:"comment,omitempty"`
	// CommitTime 报备提交时间
	CommitTime int64 `json:"commit_time,omitempty"`
	// GoodsID 商品id
	GoodsID uint64 `json:"goods_id,omitempty"`
	// MallID 店铺id
	MallID uint64 `json:"mall_id,omitempty"`
	// PromotionEndTime 推广结束时间
	PromotionEndTime int64 `json:"promotion_end_time,omitempty"`
	// PromotionStartTime 推广开始时间
	PromotionStartTime int64 `json:"promotion_start_time,omitempty"`
	// Status 报备状态。0-已创建，1-已提交，2-已通过，3-已驳回
	Status int `json:"status,omitempty"`
	// UpdatedAt 最后更新时间
	UpdatedAt int64 `json:"updated_at,omitempty"`
}

// PromotionGoodsQuery 多多进宝信息流投放商品报备进度查询
func PromotionGoodsQuery(ctx context.Context, clt *core.SDKClient, req *PromotionGoodsQueryRequest) (int, []PromotionGoodsApplication, error) {
	var resp PromotionGoodsQueryResponse
	if err := clt.Do(ctx, req, &resp, ""); err != nil {
		return 0, nil, err
	}
	return resp.Response.Total, resp.Response.List, nil
}
