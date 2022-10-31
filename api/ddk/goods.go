package ddk

import "github.com/bububa/openpdd/model"

// Goods 多多进宝商品
type Goods struct {
	// ActivityPromotionRate 活动佣金比例，千分比（特定活动期间的佣金比例）
	ActivityPromotionRate int `json:"activity_promotion_rate,omitempty" xml:"activity_promotion_rate,omitempty"`
	// ActivityTags 商品活动标记数组，例：[4,7]，4-秒杀 7-百亿补贴等
	ActivityTags []int `json:"activity_tags,omitempty" xml:"activity_tags,omitempty"`
	// BrandName 商品品牌词信息，如“苹果”、“阿迪达斯”、“李宁”等
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// CashGiftAmount 全局礼金金额，单位分
	CashGiftAmount int64 `json:"cash_gift_amount,omitempty" xml:"cash_gift_amount,omitempty"`
	// CatID 商品类目id
	CatID model.Uint64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
	// CatIDs 商品一~四级类目ID列表
	CatIDs []uint64 `json:"cat_ids,omitempty" xml:"cat_ids,omitempty"`
	// CouponDiscount 优惠券面额,单位为分
	CouponDiscount int64 `json:"coupon_discount,omitempty" xml:"coupon_discount,omitempty"`
	// CouponEnd 优惠券失效时间,UNIX时间戳
	CouponEnd int64 `json:"coupon_end,omitempty" xml:"coupon_end,omitempty"`
	// CouponMinOrderAmount 优惠券门槛价格,单位为分
	CouponMinOrderAmount int64 `json:"coupon_min_order_amount,omitempty" xml:"coupon_min_order_amount,omitempty"`
	// CouponPrice 优惠券金额
	CouponPrice int64 `json:"coupon_price,omitempty" xml:"coupon_price,omitempty"`
	// CouponRemainQuantity 优惠券剩余数量
	CouponRemainQuantity int `json:"coupon_remain_quantity,omitempty" xml:"coupon_remain_quantity,omitempty"`
	// CouponStartTime 优惠券生效时间,UNIX时间戳
	CouponStartTime int64 `json:"coupon_start_time,omitempty" xml:"coupon_start_time,omitempty"`
	// CouponTotalQuantity 优惠券总数量
	CouponTotalQuantity int `json:"coupon_total_quantity,omitempty" xml:"coupon_total_quantity,omitempty"`
	// CreateAt 创建时间
	CreateAt int64 `json:"create_at,omitempty" xml:"create_at,omitempty"`
	// DescTxt 描述分
	DescTxt string `json:"desc_txt,omitempty" xml:"desc_txt,omitempty"`
	// ExtraCouponAmount 额外优惠券，单位为分
	ExtraCouponAmount int64 `json:"extra_coupon_amount,omitempty" xml:"extra_coupon_amount,omitempty"`
	// GoodsDesc 商品描述
	GoodsDesc string `json:"goods_desc,omitempty" xml:"goods_desc,omitempty"`
	// GoodsImageURL 商品主图
	GoodsImageURL string `json:"goods_image_url,omitempty" xml:"goods_image_url,omitempty"`
	// GoodsLabels 商品特殊标签列表。例: [1]，1-APP专享
	GoodsLabels []int `json:"goods_labels,omitempty" xml:"goods_labels,omitempty"`
	// GoodsName 商品名称
	GoodsName string `json:"goods_name,omitempty" xml:"goods_name,omitempty"`
	// GoodsRate 商品等级
	GoodsRate int `json:"goods_rate,omitempty" xml:"goods_rate,omitempty"`
	// GodsID 商品ID
	GoodsID uint64 `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
	// GoodsSign 商品goodsSign，支持通过goodsSign查询商品。goodsSign是加密后的goodsId, goodsId已下线，请使用goodsSign来替代。使用说明：https://jinbao.pinduoduo.com/qa-system?questionId=252
	GoodsSign string `json:"goods_sign,omitempty" xml:"goods_sign,omitempty"`
	// GoodsThumbnailURL 商品缩略图
	GoodsThumbnailURL string `json:"goods_thumbnail_url,omitempty" xml:"goods_thumbnail_url,omitempty"`
	// GoodsType 商品类型
	GoodsType int `json:"goods_type,omitempty" xml:"goods_type,omitempty"`
	// HasCoupon 商品是否带券,true-带券,false-不带券
	HasCoupon bool `json:"has_coupon,omitempty" xml:"has_coupon,omitempty"`
	// HasMaterial 商品是否有素材(图文、视频)
	HasMaterial bool `json:"has_material,omitempty" xml:"has_material,omitempty"`
	// LgstTxt 物流分
	LgstTxt string `json:"lgst_txt,omitempty" xml:"lgst_txt,omitempty"`
	// MallID 商家id
	MallID uint64 `json:"mall_id,omitempty" xml:"mall_id,omitempty"`
	// MallName 店铺名称
	MallName string `json:"mall_name,omitempty" xml:"mall_name,omitempty"`
	// MarketFee 市场服务费
	MarketFee int64 `json:"market_fee,omitempty" xml:"market_fee,omitempty"`
	// MerchantType 商家类型
	MerchatType model.Int `json:"merchant_type,omitempty" xml:"merchant_type,omitempty"`
	// MinGroupPrice 最小成团价格，单位分
	MinGroupPrice int64 `json:"min_group_price,omitempty" xml:"min_group_price,omitempty"`
	// MinNormalPrice 最小单买价格，单位分
	MinNormalPrice int64 `json:"min_normal_price,omitempty" xml:"min_normal_price,omitempty"`
	// OptID 商品标签类目ID,使用pdd.goods.opt.get获取
	OptID model.Uint64 `json:"opt_id,omitempty" xml:"opt_id,omitempty"`
	// OptIDs 商品一~四级标签类目ID列表
	OptIDs []uint64 `json:"opt_ids,omitempty" xml:"opt_ids,omitempty"`
	// OptName 商品标签名
	OptName string `json:"opt_name,omitempty" xml:"opt_name,omitempty"`
	// PredictPromotionRate 比价行为预判定佣金，需要用户备案
	PredictPromotionRate int `json:"predict_promotion_rate,omitempty" xml:"predict_promotion_rate,omitempty"`
	// PromotionRate 佣金比例,千分比
	PromotionRate int `json:"promotion_rate,omitempty" xml:"promotion_rate,omitempty"`
	// QrcodeImageURL 二维码主图
	QrcodeImageURL string `json:"qrcode_image_url,omitempty" xml:"qrcode_image_url,omitempty"`
	// RealtimeSalesTip 商品近1小时在多多进宝的实时销量（仅实时热销榜提供）
	RealtimeSalesTip string `json:"realtime_sales_tip,omitempty" xml:"realtime_sales_tip,omitempty"`
	// SalesTip 销售量
	SalesTip string `json:"sales_tip,omitempty" xml:"sales_tip,omitempty"`
	// SearchID 搜索id，建议生成推广链接时候填写，提高收益。
	SearchID string `json:"search_id,omitempty" xml:"search_id,omitempty"`
	// ServTxt 服务分
	ServTxt string `json:"serv_txt,omitempty" xml:"serv_txt,omitempty"`
	// ShareDesc 分享描述
	ShareDesc string `json:"share_desc,omitempty" xml:"share_desc,omitempty"`
	// ShareRate 招商分成服务费比例，千分比
	ShareRate int `json:"share_rate,omitempty" xml:"share_rate,omitempty"`
	// SubsidyAmount 优势渠道专属商品补贴金额，单位为分。针对优质渠道的补贴活动，指定优势渠道可通过推广该商品获取相应补贴。补贴活动入口：[进宝网站-官方活动]
	SubsidyAmount int64 `json:"subsidy_amount,omitempty" xml:"subsidy_amount,omitempty"`
	// SubsidyDuoAmountTenMillion 官方活动给渠道的收入补贴金额，不允许直接给下级代理展示，单位为分
	SubsidyDuoAmountTenMillion int64 `json:"subsidy_duo_amount_ten_million,omitempty" xml:"subsidy_duo_amount_ten_million,omitempty"`
	// UnifiedTags 优惠标签列表，包括："X元券","比全网低X元","服务费","精选素材","近30天低价","同款低价","同款好评","同款热销","旗舰店","一降到底","招商优选","商家优选","好价再降X元","全站销量XX","实时热销榜第X名","实时好评榜第X名","额外补X元"等
	UnifiedTags []string `json:"unified_tags,omitempty" xml:"unified_tags,omitempty"`
}
