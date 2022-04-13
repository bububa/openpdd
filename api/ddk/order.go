package ddk

// Order 订单详情
type Order struct {
	// ActivityTags 商品活动标记数组，例：[4,7]，4-秒杀 7-百亿补贴等
	ActivityTags []int `json:"activity_tags,omitempty" xml:"activity_tags,omitempty"`
	// AuthDuoID 多多客工具id
	AuthDuoID uint64 `json:"auth_duo_id,omitempty" xml:"auth_duo_id,omitempty"`
	// BatchNo 结算批次号
	BathNo string `json:"batch_no,omitempty" xml:"batch_no,omitempty"`
	// CatIDs 商品一~四级类目ID列表
	CatIDs []uint64 `json:"cat_ids,omitempty" xml:"cat_ids,omitempty"`
	// CpaNew 是否是 cpa 新用户，1表示是，0表示否
	CpaNew int `json:"cpa_new,omitempty" xml:"cpa_new,omitempty"`
	// CpsSign CPS_Sign
	CpsSign string `json:"cps_sign,omitempty" xml:"cps_sign,omitempty"`
	// CustomParameters 自定义参数
	CustomParameters string `json:"custom_parameters,omitempty" xml:"custom_parameters,omitempty"`
	// FailReason 订单审核失败原因
	FailReason string `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	// GoodsCategoryName 商品一级类目名称
	GoodsCategoryName string `json:"goods_category_name,omitempty" xml:"goods_category_name,omitempty"`
	// GoodsID 商品id
	GoodsID uint64 `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
	// GoodsName 商品名称
	GoodsName string `json:"goods_name,omitempty" xml:"goods_name,omitempty"`
	// GoodsPrice 商品价格（分）
	GoodsPrice int64 `json:"goods_price,omitempty" xml:"goods_price,omitempty"`
	// GoodsQuantity 商品数量
	GoodsQuantity int64 `json:"goods_quantity,omitempty" xml:"goods_quantity,omitempty"`
	// GoodsSign goodsSign是加密后的goodsId，goodsId已下线，请使用goodsSign来替代。需要注意的是：推广链接带有goodsSign信息时，订单会返回原goodsSign；反之，会生成新的goodsSign返回。
	GoodsSign string `json:"goods_sign,omitempty" xml:"goods_sign,omitempty"`
	// GoodsThumbnailURL 商品缩略图
	GoodsThumbnailURL string `json:"goods_thumbnail_url,omitempty" xml:"goods_thumbnail_url,omitempty"`
	// GroupID 成团编号
	GroupID uint64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// IsDirect 是否直推 ，1表示是，0表示否
	IsDirect int `json:"is_direct,omitempty" xml:"is_direct,omitempty"`
	// MallID 店铺id
	MallID uint64 `json:"mall_id,omitempty" xml:"mall_id,omitempty"`
	// MallName 店铺名称
	MallName string `json:"mall_name,omitempty" xml:"mall_name,omitempty"`
	// NoSubsidyReason 非补贴订单原因，例如："商品补贴达上限"，"达到单个用户下单上限"，"非指定落地页直推订单"，"订单超过2个月未审核成功"等
	NoSubsidyReason string `json:"no_subsidy_reason,omitempty" xml:"no_subsidy_reason,omitempty"`
	// OrderAmount 订单价格（分）
	OrderAmount int64 `json:"order_amount,omitempty" xml:"order_amount,omitempty"`
	// OrderCreateTime 订单创建时间（UNIX时间戳）
	OrderCreateTime int64 `json:"order_create_time,omitempty" xml:"order_create_time,omitempty"`
	// OrderGroupSuccessTime 订单创建时间（UNIX时间戳）
	OrderGroupSuccessTime int64 `json:"order_group_success_time,omitempty" xml:"order_group_success_time,omitempty"`
	// OrderModifyAt 订单最后更新时间（UNIX时间戳）
	OrderModifyAt int64 `json:"order_modify_at,omitempty" xml:"order_modify_at,omitempty"`
	// OrderPayTime 订单支付时间（UNIX时间戳）
	OrderPayTime int64 `json:"order_pay_time,omitempty" xml:"order_pay_time,omitempty"`
	// OrderReceiveTime 订单确认收货时间（UNIX时间戳）
	OrderReceiveTime int64 `json:"order_receive_time,omitempty" xml:"order_receive_time,omitempty"`
	// OrderSettleTime 订单结算时间（UNIX时间戳）
	OrderSettleTime int64 `json:"order_settle_time,omitempty" xml:"order_settle_time,omitempty"`
	// OrderSN 订单编号
	OrderSN string `json:"order_sn,omitempty" xml:"order_sn,omitempty"`
	// OrderStatus 订单状态
	OrderStatus int `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// OrderStatusDesc 订单状态：0-已支付；1-已成团；2-确认收货；3-审核成功；4-审核失败（不可提现）；5-已经结算 ;10-已处罚
	OrderStatusDesc string `json:"order_status_desc,omitempty" xml:"order_status_desc,omitempty"`
	// OrderVerifyTime 订单审核时间（UNIX时间戳）
	OrderVerifyTime int64 `json:"order_verify_time,omitempty" xml:"order_verify_time,omitempty"`
	// Pid 推广位id
	Pid string `json:"pid,omitempty"`
	// PlatformDiscount 平台券金额，表示该订单使用的平台券金额，单位分
	PlatformDiscount int64 `json:"platform_discount,omitempty" xml:"platform_discount,omitempty"`
	// PointTime 打点时间
	PointTime int64 `json:"point_time,omitempty" xml:"point_time,omitempty"`
	// PriceCompareStatus 比价状态：0：正常，1：比价
	PriceCompareStatus int `json:"price_compare_status,omitempty" xml:"price_compare_status,omitempty"`
	// PromotionAmount 佣金（分）
	PromotionAmount int64 `json:"promotion_amount,omitempty" xml:"promotion_amount,omitempty"`
	// PromotionRate 佣金比例 千分比
	PromotionRate int `json:"promotion_rate,omitempty" xml:"promotion_rate,omitempty"`
	// RedPacketType 超级红包补贴类型：0-非红包补贴订单，1-季度新用户补贴
	RedPacketType int `json:"red_packet_type,omitempty" xml:"red_packet_type,omitempty"`
	// ReturnStatus 售后状态：0：无，1：售后中，2：售后完成
	ReturnStatus int `json:"return_status,omitempty" xml:"return_status,omitempty"`
	// SepDuoID 直播间订单推广duoId
	SepDuoID uint64 `json:"sep_duo_id,omitempty" xml:"sep_duo_id,omitempty"`
	// SepMarketFee 直播间推广佣金
	SepMarketFee int64 `json:"sep_market_fee,omitempty" xml:"sep_market_fee,omitempty"`
	// SepParams 直播间推广自定义参数
	SepParams string `json:"sep_params,omitempty" xml:"sep_params,omitempty"`
	// SepPid 直播间订单推广位
	SepPid string `json:"sep_pid,omitempty" xml:"sep_pid,omitempty"`
	// SepRate 直播间推广佣金比例
	SepRate int `json:"sep_rate,omitempty" xml:"sep_rate,omitempty"`
	// ShareAmount 招商分成服务费金额，单位为分
	ShareAmount int64 `json:"share_amount,omitempty" xml:"share_amount,omitempty"`
	// ShareRate 招商分成服务费比例，千分比
	ShareRate int `json:"share_rate,omitempty" xml:"share_rate,omitempty"`
	// SubSidyAmount 优势渠道专属商品补贴金额，单位为分。针对优质渠道的补贴活动，指定优势渠道可通过推广该商品获取相应补贴。补贴活动入口：[进宝网站-官方活动]
	SubSidyAmount int64 `json:"subsidy_amount,omitempty" xml:"subsidy_amount,omitempty"`
	// SubsidyDuoAmountLevel 等级补贴给渠道的收入补贴，不允许直接给下级代理展示，单位为分
	SubsidyDuoAmountLevel int64 `json:"subsidy_duo_amount_level,omitempty" xml:"subsidy_duo_amount_level,omitempty"`
	// SubsidyDuoAmountTenMillion 官方活动给渠道的收入补贴金额，不允许直接给下级代理展示，单位为分
	SubsidyDuoAmountTenMillion int64 `json:"subsidy_duo_amount_ten_million,omitempty" xml:"subsidy_duo_amount_ten_million,omitempty"`
	// SubsidyType 订单补贴类型：0-非补贴订单，1-千万补贴，2-社群补贴，3-多多星选，4-品牌优选，5-千万神券
	SubsidyType int `json:"subsidy_type,omitempty" xml:"subsidy_type,omitempty"`
	// Type 下单场景类型：0-单品推广，1-红包活动推广，4-多多进宝商城推广，7-今日爆款，8-品牌清仓，9-1.9包邮，77-刮刮卡活动推广，94-充值中心，101-品牌黑卡，103-百亿补贴频道，104-内购清单频道，105-超级红包
	Type int `json:"type,omitempty" xml:"type,omitempty"`
	// URLLastGenerateTime 链接最后一次生产时间
	URLLastGenerateTime int64 `json:"url_last_generate_time,omitempty" xml:"url_last_generate_time,omitempty"`
	// ZsDuoID 招商多多客id
	ZsDuoID uint64 `json:"zs_duo_id,omitempty" xml:"zs_duo_id,omitempty"`
}
