package ddk

import (
	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// OrderListIncrementGetRequest 最后更新时间段增量同步推广订单信息 API Request
type OrderListIncrementGetRequest struct {
	// CashGiftOrder 是否为礼金订单，查询礼金订单时，订单类型不填（默认推广订单）。
	CashGiftOrder bool `json:"cash_gift_order,omitempty"`
	// EndUpdateTime 查询结束时间，和开始时间相差不能超过24小时。note：此时间为时间戳，指格林威治时间 1970 年01 月 01 日 00 时 00 分 00 秒(北京时间 1970 年 01 月 01 日 08 时 00 分 00 秒)起至现在的总秒数
	EndUpdateTime int64 `json:"end_update_time,omitempty"`
	// Page 第几页，从1到10000，默认1，注：使用最后更新时间范围增量同步时，必须采用倒序的分页方式（从最后一页往回取）才能避免漏单问题。
	Page int `json:"page,omitempty"`
	// PageSize 返回的每页结果订单数，默认为100，范围为10到100，建议使用40~50，可以提高成功率，减少超时数量。
	PageSize int `json:"page_size,omitempty"`
	// QueryOrderType 订单类型：1-推广订单；2-直播间订单
	QueryOrderType int `json:"query_order_type,omitempty"`
	// ReturnCount 是否返回总数，默认为true，如果指定false, 则返回的结果中不包含总记录数，通过此种方式获取增量数据，效率在原有的基础上有80%的提升。
	ReturnCount *bool `json:"return_count,omitempty"`
	// StartUpdateTime 最近90天内多多进宝商品订单更新时间--查询时间开始。note：此时间为时间戳，指格林威治时间 1970 年01 月 01 日 00 时 00 分 00 秒(北京时间 1970 年 01 月 01 日 08 时 00 分 00 秒)起至现在的总秒数
	StartUpdateTIme int64 `json:"start_update_time,omitempty"`
}

// GetType implement Request interface
func (r OrderListIncrementGetRequest) GetType() string {
	return "pdd.ddk.order.list.increment.get"
}

// OrderListIncrementGetResponse 最后更新时间段增量同步推广订单信息 API Response
type OrderListIncrementGetResponse struct {
	model.CommonResponse
	Response struct {
		// List order list
		List []Order `json:"order_list,omitempty" xml:"order_list,omitempty"`
		// TotalCount total count
		TotalCount int `json:"total_count,omitempty" xml:"total_count,omitempty"`
	} `json:"order_list_get_response" xml:"order_list_get_response"`
}

// OrderListIncrementGet 最后更新时间段增量同步推广订单信息
func OrderListIncrementGet(clt *core.SDKClient, req *OrderListIncrementGetRequest, accessToken string) (int, []Order, error) {
	var resp OrderListIncrementGetResponse
	if err := clt.Do(req, &resp, accessToken); err != nil {
		return 0, nil, err
	}
	return resp.Response.TotalCount, resp.Response.List, nil
}
