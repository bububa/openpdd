package oauth

import (
	"github.com/bububa/openpdd/api/ddk"
	"github.com/bububa/openpdd/core"
)

// GoodsPromotionUrlGenerateRequest 多多进宝推广链接生成 API Request
type GoodsPromotionUrlGenerateRequest struct {
	// CashGiftID 多多礼金ID
	CashGiftID uint64 `json:"cash_gift_id,omitempty"`
	// CashGiftName 自定义礼金标题，用于向用户展示渠道专属福利，不超过12个字
	CashGiftName string `json:"cash_gift_name,omitempty"`
	// CustomParameters 自定义参数，为链接打上自定义标签；自定义参数最长限制64个字节；格式为： {"uid":"11111","sid":"22222"} ，其中 uid 用户唯一标识，可自行加密后传入，每个用户仅且对应一个标识，必填； sid 上下文信息标识，例如sessionId等，非必填。该json字符串中也可以加入其他自定义的key。（如果使用GET请求，请使用URLEncode处理参数）
	CustomParameters interface{} `json:"custom_parameters,omitempty"`
	// ForceDuoID 是否使用多多客专属推广计划
	ForceDuoID bool `json:"force_duo_id,omitempty"`
	// GenerateAuthorityURL 是否生成带授权的单品链接。如果未授权，则会走授权流程
	GenerateAuthorityURL bool `json:"generate_authority_url,omitempty"`
	// GenerateMallCollectCoupon 是否生成店铺收藏券推广链接
	GenerateMallCollectCoupon bool `json:"generate_mall_collect_coupon,omitempty"`
	// GenerateQQApp 是否生成qq小程序
	GenerateQQApp bool `json:"generate_qq_app,omitempty"`
	// GenerateSchemeURL 是否返回 schema URL
	GenerateSchemeURL bool `json:"generate_scheme_url,omitempty"`
	// GenerateShortURL 是否生成短链接，true-是，false-否
	GenerateShortURL bool `json:"generate_short_url,omitempty"`
	// GenerateWeApp 是否生成拼多多福利券微信小程序推广信息
	GenerateWeApp bool `json:"generate_we_app,omitempty"`
	// GoodsSignList 商品goodsSign列表，例如：["c9r2omogKFFAc7WBwvbZU1ikIb16_J3CTa8HNN"]，支持批量生链。goodsSign是加密后的goodsId, goodsId已下线，请使用goodsSign来替代。使用说明：https://jinbao.pinduoduo.com/qa-system?questionId=252
	GoodsSignList []string `json:"goods_sign_list,omitempty"`
	// MaterialID 素材ID，可以通过商品详情接口获取商品素材信息
	MaterialID string `json:"material_id,omitempty"`
	// MultiGroup true--生成多人团推广链接 false--生成单人团推广链接（默认false）1、单人团推广链接：用户访问单人团推广链接，可直接购买商品无需拼团。2、多人团推广链接：用户访问双人团推广链接开团，若用户分享给他人参团，则开团者和参团者的佣金均结算给推手
	MultiGroup bool `json:"multi_group,omitempty"`
	// Pid 推广位ID
	Pid string `json:"pid,omitempty"`
	// SearchID 搜索id，建议填写，提高收益。来自pdd.ddk.goods.recommend.get、pdd.ddk.goods.search、pdd.ddk.top.goods.list.query等接口
	SearchID string `json:"search_id,omitempty"`
	// ZsDouID 招商多多客ID
	ZsDouID uint64 `json:"zs_dou_id,omitempty"`
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
