package pdd

type Order struct {
	Address                   string             `json:"address"`
	AddressMask               string             `json:"address_mask"`
	AfterSalesStatus          int                `json:"after_sales_status"`
	BondedWarehouse           string             `json:"bonded_warehouse"`
	BuyerMemo                 string             `json:"buyer_memo"`
	CapitalFreeDiscount       float64            `json:"capital_free_discount"`
	CardInfoList              []CardInfo         `json:"card_info_list"`
	CatId1                    int64              `json:"cat_id_1"`
	CatId2                    int64              `json:"cat_id_2"`
	CatId3                    int64              `json:"cat_id_3"`
	CatId4                    int64              `json:"cat_id_4"`
	City                      string             `json:"city"`
	CityId                    int                `json:"city_id"`
	ConfirmStatus             int                `json:"confirm_status"`
	ConfirmTime               string             `json:"confirm_time"`
	ConsolidateInfo           ConsolidateInfo    `json:"consolidate_info"`
	Country                   string             `json:"country"`
	CountryId                 int                `json:"country_id"`
	CreatedTime               string             `json:"created_time"`
	DeliveryHomeValue         float64            `json:"delivery_home_value"`
	DeliveryInstallValue      float64            `json:"delivery_install_value"`
	DeliveryOneDay            int                `json:"delivery_one_day"`
	DiscountAmount            float64            `json:"discount_amount"`
	DuoduoWholesale           int                `json:"duoduo_wholesale"`
	ExtraDeliveryList         []ExtraDelivery    `json:"extra_delivery_list"`
	FreeSf                    int                `json:"free_sf"`
	GiftDeliveryList          []GiftDelivery     `json:"gift_delivery_list"`
	GiftList                  []GiftInfo         `json:"gift_list"`
	GoodsAmount               float64            `json:"goods_amount"`
	GroupOrderId              int64              `json:"group_order_id"`
	GroupRole                 int                `json:"group_role"`
	GroupStatus               int                `json:"group_status"`
	HomeDeliveryType          int                `json:"home_delivery_type"`
	HomeInstallValue          float64            `json:"home_install_value"`
	InnerTransactionId        string             `json:"inner_transaction_id加密"`
	InvoiceStatus             int                `json:"invoice_status"`
	IsLuckyFlag               int                `json:"is_lucky_flag"`
	IsPreSale                 int                `json:"is_pre_sale"`
	IsStockOut                int                `json:"is_stock_out"`
	ItemList                  []ItemInfo         `json:"item_list"`
	LastShipTime              string             `json:"last_ship_time"`
	LogisticsId               int64              `json:"logistics_id"`
	MktBizType                int                `json:"mkt_biz_type"`
	OnlySupportReplace        int                `json:"only_support_replace"`
	OrderChangeAmount         float64            `json:"order_change_amount"`
	OrderDepotInfo            DepotInfo          `json:"order_depot_info"`
	OrderSn                   string             `json:"order_sn"`
	OrderStatus               int                `json:"order_status"`
	OrderTagList              []OrderTag         `json:"order_tag_list"`
	PayAmount                 float64            `json:"pay_amount"`
	PayNo                     string             `json:"pay_no加密"`
	PayTime                   string             `json:"pay_time"`
	PayType                   string             `json:"pay_type"`
	PlatformDiscount          float64            `json:"platform_discount"`
	Postage                   float64            `json:"postage"`
	PreSaleTime               string             `json:"pre_sale_time"`
	PromiseDeliveryTime       string             `json:"promise_delivery_time"`
	Province                  string             `json:"province"`
	ProvinceId                int                `json:"province_id"`
	ReceiveTime               string             `json:"receive_time"`
	ReceiverAddress           string             `json:"receiver_address"`
	ReceiverAddressMask       string             `json:"receiver_address_mask"`
	ReceiverName              string             `json:"receiver_name"`
	ReceiverNameMask          string             `json:"receiver_name_mask"`
	ReceiverPhone             string             `json:"receiver_phone"`
	ReceiverPhoneMask         string             `json:"receiver_phone_mask"`
	RefundStatus              int                `json:"refund_status"`
	Remark                    string             `json:"remark"`
	RemarkTag                 int                `json:"remark_tag"`
	RemarkTagName             string             `json:"remark_tag_name"`
	ResendDeliveryList        []ResendDelivery   `json:"resend_delivery_list"`
	ReturnFreightPayer        int                `json:"return_freight_payer"`
	RiskControlStatus         int                `json:"risk_control_status"`
	SelfContained             int                `json:"self_contained"`
	SellerDiscount            float64            `json:"seller_discount"`
	ServiceFeeDetail          []ServiceFeeDetail `json:"service_fee_detail"`
	ShipAdditionalLinkOrder   string             `json:"ship_additional_link_order"`
	ShipAdditionalOriginOrder string             `json:"ship_additional_origin_order"`
	ShippingTime              string             `json:"shipping_time"`
	ShippingType              int                `json:"shipping_type"`
	StepOrderInfo             StepOrderInfo      `json:"step_order_info"`
	StockOutHandleStatus      int                `json:"stock_out_handle_status"`
	StoreInfo                 StoreInfo          `json:"store_info"`
	SupportNationwideWarranty int                `json:"support_nationwide_warranty"`
	Town                      string             `json:"town"`
	TownId                    int                `json:"town_id"`
	TrackingNumber            string             `json:"tracking_number"`
	TradeType                 int                `json:"trade_type"`
	UpdatedAt                 string             `json:"updated_at"`
	UrgeShippingTime          string             `json:"urge_shipping_time"`
	YypsDate                  string             `json:"yyps_date"`
	YypsTime                  string             `json:"yyps_time"`
}

type CardInfo struct {
	CardNo       string `json:"card_no"`
	MaskPassword string `json:"mask_password"`
}

type ConsolidateInfo struct {
	ConsolidateType int `json:"consolidate_type,omitempty"`
}

type ExtraDelivery struct {
	LogisticsId    int    `json:"logistics_id"`
	TrackingNumber string `json:"tracking_number"`
}

type GiftDelivery struct {
	LogisticsId    int    `json:"logistics_id"`
	TrackingNumber string `json:"tracking_number"`
}

type GiftInfo struct {
	GoodsCount   int     `json:"goods_count"`
	GoodsId      int64   `json:"goods_id"`
	GoodsImg     string  `json:"goods_img"`
	GoodsName    string  `json:"goods_name"`
	GoodsPrice   float64 `json:"goods_price"`
	GoodsSpec    string  `json:"goods_spec"`
	OuterGoodsId string  `json:"outer_goods_id"`
	OuterId      string  `json:"outer_id"`
	SkuId        int64   `json:"sku_id"`
}

type ItemInfo struct {
	GoodsCount   int     `json:"goods_count"`
	GoodsId      int64   `json:"goods_id"`
	GoodsImg     string  `json:"goods_img"`
	GoodsName    string  `json:"goods_name"`
	GoodsPrice   float64 `json:"goods_price"`
	GoodsSpec    string  `json:"goods_spec"`
	OuterGoodsId string  `json:"outer_goods_id"`
	OuterId      string  `json:"outer_id"`
	SkuId        int64   `json:"sku_id"`
}

type DepotInfo struct {
	DepotCode       string        `json:"depot_code"`
	DepotId         string        `json:"depot_id"`
	DepotName       string        `json:"depot_name"`
	DepotType       int           `json:"depot_type"`
	WareId          string        `json:"ware_id"`
	WareName        string        `json:"ware_name"`
	WareSn          string        `json:"ware_sn"`
	WareSubInfoList []WareSubInfo `json:"ware_sub_info_list"`
}

type WareSubInfo struct {
	WareId       int64  `json:"ware_id"`
	WareName     string `json:"ware_name"`
	WareQuantity int64  `json:"ware_quantity"`
	WareSn       string `json:"ware_sn"`
	WareType     int    `json:"ware_type"`
}

type OrderTag struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type ResendDelivery struct {
	LogisticsId    int    `json:"logistics_id"`
	TrackingNumber string `json:"tracking_number"`
}

type ServiceFeeDetail struct {
	ServiceFee  float64 `json:"service_fee"`
	ServiceName string  `json:"service_name"`
}

type StepOrderInfo struct {
	AdvancedPaidFee    float64 `json:"advanced_paid_fee"`
	StepDiscountAmount float64 `json:"step_discount_amount"`
	StepPaidFee        float64 `json:"step_paid_fee"`
	StepTradeStatus    int     `json:"step_trade_status"`
}

type StoreInfo struct {
	StoreId     int64  `json:"store_id"`
	StoreName   string `json:"store_name"`
	StoreNumber string `json:"store_number"`
}
