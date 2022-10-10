package ddk

import (
	"errors"

	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// GoodsPromotionRightAuthRequest 多多进宝信息流渠道备案授权素材上传接口 API Request
type GoodsPromotionRightAuthRequest struct {
	// DemoURL 推广商品视频素材url
	DemoURL string `json:"demo_url,omitempty"`
	// DuoID 渠道duoId
	DuoID uint64 `json:"duo_id,omitempty"`
	// GoodsID 商品GoodsId
	GoodsID uint64 `json:"goods_id,omitempty"`
	// MallCertificationURL 商家资质证明图片url列表，1到3张图
	MallCertificationURL []string `json:"mall_certificate_url,omitempty"`
	// PromotionCodeURL 推广视频预览码url
	PromotionCodeURL string `json:"promotion_code_url,omitempty"`
	// PromotionEndTime 推广结束时间戳，毫秒
	PromotionEndTime int64 `json:"promotion_end_time,omitempty"`
	// PromotionStartTime 推广开始时间戳，毫秒
	PromotionStartTime int64 `json:"promotion_start_time,omitempty"`
	// ThumbPicURL 商品图片素材url列表，0到3张图
	ThumbPicURL []string `json:"thumb_pic_url,omitempty"`
}

// GetType implement Request interface
func (r GoodsPromotionRightAuthRequest) GetType() string {
	return "pdd.ddk.goods.promotion.right.auth"
}

// GoodsPromotionRightAuthResponse 多多进宝信息流渠道备案授权素材上传接口 API Response
type GoodsPromotionRightAuthResponse struct {
	model.CommonResponse
	Response struct {
		// Reason 备案失败原因
		Reason string `json:"reason,omitempty"`
		// Result 备案结果
		Result bool `json:"result,omitempty"`
	} `json:"goods_promotion_right_auth_response"`
}

// GoodsPromotionRightAuth 多多进宝信息流渠道备案授权素材上传接口
func GoodsPromotionRightAuth(clt *core.SDKClient, req *GoodsPromotionRightAuthRequest) error {
	var resp GoodsPromotionRightAuthResponse
	if err := clt.Do(req, &resp, ""); err != nil {
		return err
	}
	if !resp.Response.Result {
		return errors.New(resp.Response.Reason)
	}
	return nil
}
