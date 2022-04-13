package oauth

import (
	"github.com/bububa/openpdd/api/ddk"
	"github.com/bububa/openpdd/core"
)

// GoodsDetailRequest 多多进宝商品详情查询 API Request
type GoodsDetailRequest struct {
	// CustomParameters 自定义参数，为链接打上自定义标签；自定义参数最长限制64个字节；格式为： {"uid":"11111","sid":"22222"} ，其中 uid 用户唯一标识，可自行加密后传入，每个用户仅且对应一个标识，必填； sid 上下文信息标识，例如sessionId等，非必填。该json字符串中也可以加入其他自定义的key。（如果使用GET请求，请使用URLEncode处理参数）
	CustomParameters interface{} `json:"custom_parameters,omitempty"`
	// GoodsImageType 商品主图类型：1-场景图，2-白底图，默认为0
	GoodsImageType int `json:"goods_image_type,omitempty"`
	// GoodsSign 商品goodsSign，支持通过goodsSign查询商品。goodsSign是加密后的goodsId, goodsId已下线，请使用goodsSign来替代。使用说明：https://jinbao.pinduoduo.com/qa-system?questionId=252
	GoodsSign string `json:"goods_sign,omitempty"`
	// NeedSkuInfo 是否获取sku信息，默认false不返回。（特殊渠道权限，需额外申请）
	NeedSkuInfo bool `json:"need_sku_info,omitempty"`
	// Pid 推广位id
	Pid string `json:"pid,omitempty"`
	// SearchID 搜索id，建议填写，提高收益。来自pdd.ddk.goods.recommend.get、pdd.ddk.goods.search、pdd.ddk.top.goods.list.query等接口
	SearchID string `json:"search_id,omitempty"`
	// ZsDuoID 招商多多客ID
	ZsDuoID uint64 `json:"zs_duo_id,omitempty"`
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
