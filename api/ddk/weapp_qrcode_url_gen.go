package ddk

import (
	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// WeappQrcodeUrlGenRequest 多多客生成单品推广小程序二维码url API Request
type WeappQrcodeUrlGenRequest struct {
	// CashGiftID 多多礼金ID
	CashGiftID uint64 `json:"cash_gift_id,omitempty"`
	// CustomParameters 自定义参数，为链接打上自定义标签；自定义参数最长限制64个字节；格式为： {"uid":"11111","sid":"22222"} ，其中 uid 用户唯一标识，可自行加密后传入，每个用户仅且对应一个标识，必填； sid 上下文信息标识，例如sessionId等，非必填。该json字符串中也可以加入其他自定义的key。（如果使用GET请求，请使用URLEncode处理参数）
	CustomParameters interface{} `json:"custom_parameters,omitempty"`
	// GenerateMallCollectCoupon 是否生成店铺收藏券推广链接
	GenerateMallCollectCoupon bool `json:"generate_mall_collect_coupon,omitempty"`
	// GoodsSignList 商品goodsSign列表，例如：["c9r2omogKFFAc7WBwvbZU1ikIb16_J3CTa8HNN"]，支持批量生链。goodsSign是加密后的goodsId, goodsId已下线，请使用goodsSign来替代。使用说明：https://jinbao.pinduoduo.com/qa-system?questionId=252
	GoodsSignList []string `json:"goods_sign_list,omitempty"`
	// Pid 推广位ID
	Pid string `json:"p_id,omitempty"`
	// ZsDouID 招商多多客ID
	ZsDouID uint64 `json:"zs_dou_id,omitempty"`
}

// GetType implement Request interface
func (r WeappQrcodeUrlGenRequest) GetType() string {
	return "pdd.ddk.weapp.qrcode.url.gen"
}

// WeappQrcodeUrlGenResponse 多多客生成单品推广小程序二维码url API Response
type WeappQrcodeUrlGenResponse struct {
	model.CommonResponse
	Response struct {
		// URL 单品推广小程序二维码url
		URL string `json:"url" xml:"url"`
	} `json:"weapp_qrcode_generate_response" xml:"weapp_qrcode_generate_response"`
}

// WeappQrcodeUrlGen 多多客生成单品推广小程序二维码url
func WeappQrcodeUrlGen(clt *core.SDKClient, req *WeappQrcodeUrlGenRequest) (string, error) {
	var resp WeappQrcodeUrlGenResponse
	if err := clt.Do(req, &resp, ""); err != nil {
		return "", err
	}
	return resp.Response.URL, nil
}
