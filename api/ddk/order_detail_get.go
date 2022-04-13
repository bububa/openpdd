package ddk

import (
	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// OrderDetailGetRequest 查询订单详情 API Request
type OrderDetailGetRequest struct {
	// OrderSN 订单号
	OrderSN string `json:"order_sn,omitempty"`
	// QueryOrderType 订单类型：1-推广订单；2-直播间订单
	QueryOrderType int `json:"query_order_type,omitempty"`
}

// GetType implement Request interface
func (r OrderDetailGetRequest) GetType() string {
	return "pdd.ddk.order.detail.get"
}

// OrderDetailGetResponse 查询订单详情 API Response
type OrderDetailGetResponse struct {
	model.CommonResponse
	Response *Order `json:"order_detail_response,omitempty" xml:"order_detail_response,omitempty"`
}

// OrderDetailGet 查询订单详情
func OrderDetailGet(clt *core.SDKClient, req *OrderDetailGetRequest, accessToken string) (*Order, error) {
	var resp OrderDetailGetResponse
	if err := clt.Do(req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Response, nil
}
