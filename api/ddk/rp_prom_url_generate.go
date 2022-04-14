package ddk

import (
	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// RpPromUrlGenerateRequest 生成营销工具推广链接 API Request
type RpPromUrlGenerateRequest struct {
	// Amount 初始金额（单位分），有效金额枚举值：300、500、700、1100和1600，默认300
	Amount int64 `json:"amount,omitempty"`
	// ChannelType 营销工具类型，必填：-1-活动列表，0-红包(需申请推广权限)，2–新人红包，3-刮刮卡，5-员工内购，10-生成绑定备案链接，12-砸金蛋，14-千万补贴B端页面，15-充值中心B端页面，16-千万补贴C端页面，17-千万补贴投票页面，23-超级红包，24-礼金全场N折活动B端页面，27-带货赢千万，28-满减券活动B端页面，29-满减券活动C端页面，30-免单B端页面，31-免单C端页面，32-转盘得现金B端页面，33-转盘得现金C端页面，34-千万神券C端页面，35-千万神券B端页面；红包推广权限申请流程链接：https://jinbao.pinduoduo.com/qa-system?questionId=289
	ChannelType int `json:"channel_type,omitempty"`
	// CustomParameters 自定义参数，为链接打上自定义标签；自定义参数最长限制64个字节；格式为： {"uid":"11111","sid":"22222"} ，其中 uid 用户唯一标识，可自行加密后传入，每个用户仅且对应一个标识，必填； sid 上下文信息标识，例如sessionId等，非必填。该json字符串中也可以加入其他自定义的key。（如果使用GET请求，请使用URLEncode处理参数）
	CustomParameters interface{} `json:"custom_parameters,omitempty"`
	// DiyYuanParam 一元购自定义参数，json格式，例如:{"goods_sign":"Y9b2_0uSWMFPGSaVwvfZAlm_y2ADLWZl_JQ7UYaS80K"}
	DiyYuanParam *DiyYuanParam `json:"diy_yuan_param,omitempty"`
	// DiyRedPacketParam 红包自定义参数，json格式
	DiyRedPacketParam *DiyRedPacketParam `json:"diy_red_packet_param,omitempty"`
	// GenerateQQApp 是否生成qq小程序
	GenerateQQApp bool `json:"generate_qq_app,omitempty"`
	// GenerateSchemeURL 是否返回 schema URL
	GenerateSchemeURL bool `json:"generate_scheme_url,omitempty"`
	// GenerateShortURL 是否生成短链接。true-是，false-否，默认false
	GenerateShortURL bool `json:"generate_short_url,omitempty"`
	// GenerateWeApp 是否生成拼多多福利券微信小程序推广信息
	GenerateWeApp bool `json:"generate_we_app,omitempty"`
	// PidList 推广位列表，长度最大为1，例如：["60005_612"]。活动页生链要求传入授权备案信息，不支持批量生链。
	PidList []string `json:"p_id_list,omitempty"`
	// ScratchCardAmount 刮刮卡指定金额（单位分），可指定2-100元间数值，即有效区间为：[200,10000]
	ScratchCardAmount int64 `json:"scratch_card_amount,omitempty"`
}

// DiyYuanParam 一元购自定义参数，json格式，例如:{"goods_sign":"Y9b2_0uSWMFPGSaVwvfZAlm_y2ADLWZl_JQ7UYaS80K"}
type DiyYuanParam struct {
	// GoodsSign 商品goodsSign，支持通过goodsSign查询商品。goodsSign是加密后的goodsId, goodsId已下线，请使用goodsSign来替代。使用说明：https://jinbao.pinduoduo.com/qa-system?questionId=252
	GoodsSign string `json:"goods_sign,omitempty"`
}

// DiyRedPacketParam 红包自定义参数，json格式
type DiyRedPacketParam struct {
	// AmountProbability 红包金额列表，200、300、500、1000、2000，单位分。红包金额和红包抵后价设置只能二选一，默认设置了红包金额会忽略红包抵后价设置
	AmountProbability []int64 `json:"amount_probability,omitempty"`
	// DisText 设置玩法，false-现金红包, true-现金券
	DisText *bool `json:"dis_text,omitempty"`
	// NotShowBackground 推广页设置，false-红包开启页, true-红包领取页
	NotShowBackground *bool `json:"not_show_background,omitempty"`
	// OptID 优先展示类目
	OptID uint64 `json:"opt_id,omitempty"`
	// RangeItems 自定义红包抵后价和商品佣金区间对象数组
	RangeItems []RangeItem `json:"range_items,omitempty"`
}

// GetType implement Request interface
func (r RpPromUrlGenerateRequest) GetType() string {
	return "pdd.ddk.rp.prom.url.generate"
}

// RpPromUrlGenerateResponse 生成营销工具推广链接 API Response
type RpPromUrlGenerateResponse struct {
	model.CommonResponse
	Response *RpPromUrlGenerateResult `json:"rp_promotion_url_generate_response,omitempty" xml:"rp_promotion_url_generate_response,omitempty"`
}

// RpPromUrlGenerateResult .
type RpPromUrlGenerateResult struct {
	// URLList url_list
	URLList []PromURL `json:"url_list,omitempty" xml:"url_list,omitempty"`
	// ResourceList resource_list
	ResourceList []PromResource `json:"resource_list,omitempty" xml:"resource_list,omitempty"`
}

// RpPromUrlGenerate 生成营销工具推广链接
func RpPromUrlGenerate(clt *core.SDKClient, req *RpPromUrlGenerateRequest) (*RpPromUrlGenerateResult, error) {
	var resp RpPromUrlGenerateResponse
	if err := clt.Do(req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Response, nil
}
