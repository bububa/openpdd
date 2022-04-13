package oauth

import (
	"github.com/bububa/openpdd/api/ddk"
	"github.com/bububa/openpdd/core"
)

// GoodsSearchRequest 多多进宝商品查询 API Request
type GoodsSearchRequest struct {
	// ActivityTags 活动商品标记数组，例：[4,7]，4-秒杀，7-百亿补贴，10851-千万补贴，11879-千万神券，10913-招商礼金商品，31-品牌黑标，10564-精选爆品-官方直推爆款，10584-精选爆品-团长推荐，24-品牌高佣，其他的值请忽略
	ActivityTags []int `json:"activity_tags,omitempty"`
	// BlockCatPackages 屏蔽商品类目包：1-拼多多小程序屏蔽的类目&关键词;2-虚拟类目;3-医疗器械;4-处方药;5-非处方药
	BlockCatPackages []uint64 `json:"block_cat_packages,omitempty"`
	// BlockCats 自定义屏蔽一级/二级/三级类目ID，自定义数量不超过20个;使用pdd.goods.cats.get接口获取cat_id
	BlockCats []uint64 `json:"block_cats,omitempty"`
	// CatID 商品类目ID，使用pdd.goods.cats.get接口获取
	CatID uint64 `json:"cat_id,omitempty"`
	// CustomParameters 自定义参数，为链接打上自定义标签；自定义参数最长限制64个字节；格式为： {"uid":"11111","sid":"22222"} ，其中 uid 用户唯一标识，可自行加密后传入，每个用户仅且对应一个标识，必填； sid 上下文信息标识，例如sessionId等，非必填。该json字符串中也可以加入其他自定义的key。（如果使用GET请求，请使用URLEncode处理参数）
	CustomParameters interface{} `json:"custom_parameters,omitempty"`
	// ForceAuthDuoID 是否使用工具商专属推广计划，默认为false
	ForceAuthDuoID bool `json:"force_auth_duo_id,omitempty"`
	// GoodsImageType 商品主图类型：1-场景图，2-白底图，默认为0
	GoodsImageType int `json:"goods_image_type,omitempty"`
	// GoodsSignList 商品goodsSign列表，例如：["c9r2omogKFFAc7WBwvbZU1ikIb16_J3CTa8HNN"]，支持通过goodsSign查询商品。goodsSign是加密后的goodsId, goodsId已下线，请使用goodsSign来替代。使用说明：https://jinbao.pinduoduo.com/qa-system?questionId=252
	GoodsSignList []string `json:"goods_sign_list,omitempty"`
	// IsBrandGoods 是否为品牌商品
	IsBrandGoods bool `json:"is_brand_goods,omitempty"`
	// Keyword 商品关键词，与opt_id字段选填一个或全部填写。可支持goods_id、拼多多链接（即拼多多app商详的链接）、进宝长链/短链（即为pdd.ddk.goods.promotion.url.generate接口生成的长短链）
	Keyword string `json:"keyword,omitempty"`
	// ListID 翻页时建议填写前页返回的list_id值
	ListID string `json:"list_id,omitempty"`
	// MerchantType 店铺类型，1-个人，2-企业，3-旗舰店，4-专卖店，5-专营店，6-普通店（未传为全部）
	MerchantType int `json:"merchant_type,omitempty"`
	// MerchantTypeList 店铺类型数组，例如：[1,2]
	MerchantTypeList []int `json:"merchant_type_list,omitempty"`
	// OptID 商品标签类目ID，使用pdd.goods.opt.get获取
	OptID uint64 `json:"opt_id,omitempty"`
	// Page 默认值1，商品分页数
	Page int `json:"page,omitempty"`
	// PageSize 默认100，每页商品数量
	PageSize int `json:"page_size,omitempty"`
	// Pid 推广位id
	Pid string `json:"pid,omitempty"`
	// RangeList 筛选范围列表 样例：[{"range_id":0,"range_from":1,"range_to":1500},{"range_id":1,"range_from":1,"range_to":1500}]
	RangeList []ddk.RangeItem `json:"range_list,omitempty"`
	// SortType 排序方式:0-综合排序;1-按佣金比率升序;2-按佣金比例降序;3-按价格升序;4-按价格降序;5-按销量升序;6-按销量降序;7-优惠券金额排序升序;8-优惠券金额排序降序;9-券后价升序排序;10-券后价降序排序;11-按照加入多多进宝时间升序;12-按照加入多多进宝时间降序;13-按佣金金额升序排序;14-按佣金金额降序排序;15-店铺描述评分升序;16-店铺描述评分降序;17-店铺物流评分升序;18-店铺物流评分降序;19-店铺服务评分升序;20-店铺服务评分降序;27-描述评分击败同类店铺百分比升序，28-描述评分击败同类店铺百分比降序，29-物流评分击败同类店铺百分比升序，30-物流评分击败同类店铺百分比降序，31-服务评分击败同类店铺百分比升序，32-服务评分击败同类店铺百分比降序
	SortType int `json:"sort_type,omitempty"`
	// UseCustomized 是否使用个性化推荐，true表示使用，false表示不使用，默认true。
	UseCustomized *bool `json:"use_customized,omitempty"`
	// WithCoupon 是否只返回优惠券的商品，false返回所有商品，true只返回有优惠券的商品
	WithCoupon bool `json:"with_coupon,omitempty"`
}

// GetType implement Request interface
func (r GoodsSearchRequest) GetType() string {
	return "pdd.ddk.oauth.goods.search"
}

// GoodsSearch 多多进宝商品查询
func GoodsSearch(clt *core.SDKClient, req *GoodsSearchRequest, accessToken string) (*ddk.GoodsSearchResult, error) {
	var resp ddk.GoodsSearchResponse
	if err := clt.Do(req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Response, nil
}
