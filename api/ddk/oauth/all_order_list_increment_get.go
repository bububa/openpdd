package oauth

import (
	"github.com/bububa/openpdd/api/ddk"
	"github.com/bububa/openpdd/core"
)

// AllOrderListIncrementGetRequest 查询所有授权的多多客订单API Request
type AllOrderListIncrementGetRequest struct {
	// EndUpdateTime 查询结束时间，和开始时间相差不能超过24小时。note：此时间为时间戳，指格林威治时间 1970 年01 月 01 日 00 时 00 分 00 秒(北京时间 1970 年 01 月 01 日 08 时 00 分 00 秒)起至现在的总秒数
	EndUpdateTime int64 `json:"end_update_time,omitempty"`
	// Page 第几页，从1到10000，默认1，注：使用最后更新时间范围增量同步时，必须采用倒序的分页方式（从最后一页往回取）才能避免漏单问题。
	Page int `json:"page,omitempty"`
	// PageSize 返回的每页结果订单数，默认为100，范围为10到100，建议使用40~50，可以提高成功率，减少超时数量。
	PageSize int `json:"page_size,omitempty"`
	// QueryOrderType 订单类型：1-推广订单；2-直播间订单
	QueryOrderType int `json:"query_order_type,omitempty"`
	// StartUpdateTime 最近90天内多多进宝商品订单更新时间--查询时间开始。note：此时间为时间戳，指格林威治时间 1970 年01 月 01 日 00 时 00 分 00 秒(北京时间 1970 年 01 月 01 日 08 时 00 分 00 秒)起至现在的总秒数
	StartUpdateTIme int64 `json:"start_update_time,omitempty"`
}

// GetType implement Request interface
func (r AllOrderListIncrementGetRequest) GetType() string {
	return "pdd.ddk.all.order.list.increment.get"
}

// AllOrderListIncrementGet 查询所有授权的多多客订单
func AllOrderListIncrementGet(clt *core.SDKClient, req *AllOrderListIncrementGetRequest, accessToken string) (int, []ddk.Order, error) {
	var resp ddk.OrderListIncrementGetResponse
	if err := clt.Do(req, &resp, accessToken); err != nil {
		return 0, nil, err
	}
	return resp.Response.TotalCount, resp.Response.List, nil
}
