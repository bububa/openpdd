package ddk

import (
	"context"

	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// CashgiftCreateRequest 创建多多礼金 API Request
type CashgiftCreateRequest struct {
	// AcquireEndTime 券批次领取结束时间。note：此时间为时间戳，指格林威治时间 1970 年01 月 01 日 00 时 00 分 00 秒(北京时间 1970 年 01 月 01 日 08 时 00 分 00 秒)起至现在的总秒数
	AcquireEndTime int64 `json:"acquire_end_time,omitempty"`
	// AcquireStartTime 券批次领取开始时间。note：此时间为时间戳，指格林威治时间 1970 年01 月 01 日 00 时 00 分 00 秒(北京时间 1970 年 01 月 01 日 08 时 00 分 00 秒)起至现在的总秒数
	AcquireStartTime int64 `json:"acquire_start_time,omitempty"`
	// AutoTake 是否自动领券，默认false不自动领券
	AutoTake bool `json:"auto_take,omitempty"`
	// CouponAmount 礼金券面额，单位为分，创建固定面额礼金券必填（创建灵活面额礼金券时，券面额 = 商品券后价 - 期望礼金券后价）
	CouponAmount int64 `json:"coupon_amount,omitempty"`
	// CouponThresholdAmount 满减门槛，单位为分。满减门槛至少需为礼金券面额的2倍，仅对固定面额礼金券生效。
	CouponThresholdAmount int64 `json:"coupon_threshold_amount,omitempty"`
	// Duration 活动持续时间，validity_time_type为 1 时必填。相对时间类型为天级时，最大值为30，即领取后30天内有效；相对时间类型为小时级时，最大值为24，即领取后24小时内有效；相对时间类型为分钟级时，则最大值为60，即领取后60分钟内有效。
	Duration int `json:"duration,omitempty"`
	// ExpectAmount 期望礼金券后价，单位为分，最小值为1。创建灵活券 (generate_flexible_coupon为true) 时必填
	ExpectAmount int64 `json:"expect_amount,omitempty"`
	// FetchRiskCheck 领券是否过风控，默认true为过风控。
	FetchRiskCheck *bool `json:"fetch_risk_check,omitempty"`
	// FreezeCondition 收益保护开关开启(rate_decrease_monitor = true)时必填。0-监控项发生降低；1-监控项低于礼金面额，默认为0。
	FreezeCondition int `json:"freeze_condition,omitempty"`
	// FreezeWatchType 收益保护开关开启(rate_decrease_monitor = true)时必填。0-佣金；1-补贴；2-佣金+补贴，默认为0。
	FreezeWatchType int `json:"freeze_watch_type,omitempty"`
	// GenerateFlexibleCoupon 是否为灵活面额礼金券，默认false为固定面额礼金券
	GenerateFlexibleCoupon bool `json:"generate_flexible_coupon,omitempty"`
	// GenerateGlobal 是否开启全场景推广，默认false不开启全场景推广，仅支持固定面额且限定商品的礼金活动。
	GenerateGlobal bool `json:"generate_global,omitempty"`
	// GoodsSignList 商品goodsSign列表，例如：["c9r2omogKFFAc7WBwvbZU1ikIb16_J3CTa8HNN"]，最多可支持传20个商品；若传空，则为不限商品礼金，不支持创建不限商品灵活礼金。goodsSign使用说明：https://jinbao.pinduoduo.com/qa-system?questionId=252
	GoodsSignList []string `json:"goods_sign_list,omitempty"`
	// LinkAcquireLimit 活动单链接可领券数量，默认无限制，最小值为1。
	LinkAcquireLimit int `json:"link_acquire_limit,omitempty"`
	// Name 礼金名称
	Name string `json:"name,omitempty"`
	// PidList 可使用推广位列表，例如：["60005_612"]。(列表中的PID方可推广该礼金)
	PidList []string `json:"p_id_list,omitempty"`
	// Quantity 礼金券数量，创建固定面额礼金券必填（创建灵活面额礼金券时，礼金券数量不固定，礼金总预算用完为止）
	Quantity int64 `json:"quantity,omitempty"`
	// RateDecreaseMonitor 收益保护开关，默认false表示关闭，仅支持固定面额且限定商品的礼金活动。开启状态下，系统将根据设置内容进行监控，当监控项满足冻结条件时，系统自动冻结礼金暂停推广，防止资金损失。（可通过多多礼金状态更新接口自行恢复推广）
	RateDecreaseMonitor bool `json:"rate_decrease_monitor,omitempty"`
	// RelativeTimeType 相对时间类型：1-天级；2-小时级；3-分钟级，有效期类型validityTimeType = 1时必填，默认为1。 例如: relative_time_type = 2, duration = 15, 表示领取后15小时内有效。
	RelativeTimeType int `json:"relative_time_type,omitempty"`
	// TotalAmount 礼金总预算，单位为分，创建灵活券 (generate_flexible_coupon为true) 时必填。默认情况，总金额 = 礼金券数量 * 礼金券面额
	TotalAmount int64 `json:"total_amount,omitempty"`
	// UserLimit 单用户可领券数量，可设置范围为1~10张，默认为1张。
	UserLimit int `json:"user_limit,omitempty"`
	// ValidityEndTime 券批次使用结束时间, validity_time_type为 2 时必填。note：此时间为时间戳，指格林威治时间 1970 年01 月 01 日 00 时 00 分 00 秒(北京时间 1970 年 01 月 01 日 08 时 00 分 00 秒)起至现在的总秒数
	ValidityEndTime int64 `json:"validity_end_time,omitempty"`
	// ValidityStartTime 券批次使用开始时间, validity_time_type为 2 时必填。note：此时间为时间戳，指格林威治时间 1970 年01 月 01 日 00 时 00 分 00 秒(北京时间 1970 年 01 月 01 日 08 时 00 分 00 秒)起至现在的总秒数
	ValidityStartTime int64 `json:"validity_start_time,omitempty"`
	// ValidityTimeType 有效期类型：1-领取后几天内有效；2-固定时间内有效
	ValidityTimeType int `json:"validity_time_type,omitempty"`
}

// GetType implement Request interface
func (r CashgiftCreateRequest) GetType() string {
	return "pdd.ddk.cashgift.create"
}

// CashgiftCreateResponse 创建多多礼金 API Response
type CashgiftCreateResponse struct {
	model.CommonResponse
	Response *CashgiftCreateResult `json:"create_cashgift_response,omitempty" xml:"create_cashgift_response,omitempty"`
}

type CashgiftCreateResult struct {
	// CashgiftID 礼金ID
	CashgiftID uint64 `json:"cash_gift_id,omitempty"`
	// Success 创建结果
	Success bool `json:"success,omitempty"`
}

// CashgiftCreate 创建多多礼金
func CashgiftCreate(ctx context.Context, clt *core.SDKClient, req *CashgiftCreateRequest) (*CashgiftCreateResult, error) {
	var resp CashgiftCreateResponse
	if err := clt.Do(ctx, req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Response, nil
}
