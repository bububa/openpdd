package ddk

import (
	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// OrderListRangeGetRequest 用时间段查询推广订单接口 API Request
type OrderListRangeGetRequest struct {
	// CashGiftOrder 是否为礼金订单，查询礼金订单时，订单类型不填（默认推广订单）。
	CashGiftOrder bool `json:"cash_gift_order,omitempty"`
	// EndTime 支付结束时间，格式: "yyyy-MM-dd HH:mm:ss" ，比如 "2020-12-01 00:00:00"
	EndTime string `json:"end_time,omitempty"`
	// StartTime 支付起始时间，格式: "yyyy-MM-dd HH:mm:ss" ，比如 "2020-12-01 00:00:00"
	StartTime string `json:"start_time,omitempty"`
	// LastOrderID 上一次的迭代器id(第一次不填)
	LastOrderID string `json:"last_order_id,omitempty"`
	// PageSize 每次请求多少条，建议300
	PageSize int `json:"page_size,omitempty"`
	// QueryOrderType 订单类型：1-推广订单；2-直播间订单
	QueryOrderType int `json:"query_order_type,omitempty"`
}

// GetType implement Request interface
func (r OrderListRangeGetRequest) GetType() string {
	return "pdd.ddk.order.list.range.get"
}

// OrderListRangeGetResponse 用时间段查询推广订单接口 API Response
type OrderListRangeGetResponse struct {
	model.CommonResponse
	Response struct {
		// LastOrderID last_order_id
		LastOrderID string `json:"last_order_id,omitempty" xml:"last_order_id,omitempty"`
		// OrderList 多多进宝推广位对象列表
		OrderList []Order `json:"order_list,omitempty" xml:"order_list,omitempty"`
	} `json:"order_list_get_response" xml:"order_list_get_response"`
}

// OrderListRangeGet 用时间段查询推广订单接口
func OrderListRangeGet(clt *core.SDKClient, req *OrderListRangeGetRequest) (string, []Order, error) {
	var resp OrderListRangeGetResponse
	if err := clt.Do(req, &resp, ""); err != nil {
		return "", nil, err
	}
	return resp.Response.LastOrderID, resp.Response.OrderList, nil
}
