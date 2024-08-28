package ddk

import (
	"context"

	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// GoodsRecommendGetRequest 多多进宝商品推荐API API Request
type GoodsRecommendGetRequest struct {
	// ActivityTags 活动商品标记数组，例：[4,7]，4-秒杀，7-百亿补贴，10851-千万补贴，11879-千万神券，10913-招商礼金商品，31-品牌黑标，10564-精选爆品-官方直推爆款，10584-精选爆品-团长推荐，24-品牌高佣，其他的值请忽略
	ActivityTags []int `json:"activity_tags,omitempty"`
	// CatID 猜你喜欢场景的商品类目，20100-百货，20200-母婴，20300-食品，20400-女装，20500-电器，20600-鞋包，20700-内衣，20800-美妆，20900-男装，21000-水果，21100-家纺，21200-文具,21300-运动,21400-虚拟,21500-汽车,21600-家装,21700-家具,21800-医药;
	CatID uint64 `json:"cat_id,omitempty"`
	// ChannelType 进宝频道推广商品: 1-今日销量榜,3-相似商品推荐,4-猜你喜欢(和进宝网站精选一致),5-实时热销榜,6-实时收益榜。默认值5
	ChannelType int `json:"channel_type,omitempty"`
	// CustomParameters 自定义参数，为链接打上自定义标签；自定义参数最长限制64个字节；格式为： {"uid":"11111","sid":"22222"} ，其中 uid 为用户唯一标识，可自行加密后传入，每个用户仅且对应一个标识，必填； sid 为上下文信息标识，例如sessionId等，非必填。该json字符串中也可以加入其他自定义的key。
	CustomParameters interface{} `json:"custom_parameters,omitempty"`
	// GoodsImageType 商品主图类型：1-场景图，2-白底图，默认为0
	GoodsImageType int `json:"goods_image_type,omitempty"`
	// GoodsSignList 商品goodsSign列表，相似商品推荐场景时必传，仅取数组的第一位，例如：["c9r2omogKFFAc7WBwvbZU1ikIb16_J3CTa8HNN"]。goodsSign是加密后的goodsId, goodsId已下线，请使用goodsSign来替代。使用说明：https://jinbao.pinduoduo.com/qa-system?questionId=252
	GoodsSignList []string `json:"goods_sign_list,omitempty"`
	// Limit 一页请求数量；默认值 ： 20
	Limit int `json:"limit,omitempty"`
	// ListID 翻页时建议填写前页返回的list_id值
	ListID string `json:"list_id,omitempty"`
	// Offset 从多少位置开始请求；默认值 ： 0，offset需是limit的整数倍，仅支持整页翻页
	Offest int `json:"offset,omitempty"`
	// Pid 推广位id
	Pid string `json:"pid,omitempty"`
}

// GetType implement Request interface
func (r GoodsRecommendGetRequest) GetType() string {
	return "pdd.ddk.goods.recommend.get"
}

// GoodsRecommendGetResponse 多多进宝商品推荐API API Response
type GoodsRecommendGetResponse struct {
	model.CommonResponse
	Response *GoodsRecommendGetResult `json:"goods_basic_detail_response,omitempty" xml:"goods_basic_detail_response,omitempty"`
}

// GoodsRecommendGetResult .
type GoodsRecommendGetResult struct {
	// List 列表
	List []Goods `json:"list,omitempty" xml:"list,omitempty"`
	// ListID 翻页时必填前页返回的list_id值
	ListID string `json:"list_id,omitempty" xml:"list_id,omitempty"`
	// SearchID 搜索id，建议生成推广链接时候填写，提高收益。
	SearchID string `json:"search_id,omitempty" xml:"search_id,omitempty"`
	// Total total
	Total int `json:"total,omitempty" xml:"total,omitempty"`
}

// GoodsRecommendGet 多多进宝商品推荐API
func GoodsRecommendGet(ctx context.Context, clt *core.SDKClient, req *GoodsRecommendGetRequest) (*GoodsRecommendGetResult, error) {
	var resp GoodsRecommendGetResponse
	if err := clt.Do(ctx, req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Response, nil
}
