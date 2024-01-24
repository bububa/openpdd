package pdd

import (
	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// OrderDetailGetRequest 查询订单详情 API Request
type OrderDetailGetRequest struct {
	// OrderSN 订单号
	OrderSN string `json:"order_sn,omitempty"`
}

// GetType implement Request interface
func (r OrderDetailGetRequest) GetType() string {
	return "pdd.order.information.get"
}

// OrderInformationGetResponse 查询订单详情 API Response
type OrderInformationGetResponse struct {
	model.CommonResponse
	Response *OrderInfoGetResponse `json:"order_info_get_response,omitempty" xml:"order_info_get_response,omitempty"`
}

type OrderInfoGetResponse struct {
	OrderInfo *Order `json:"order_info,omitempty" xml:"order_info,omitempty"`
}

// OrderDetailGet 查询订单详情
func OrderDetailGet(clt *core.SDKClient, req *OrderDetailGetRequest) (*Order, error) {
	var resp OrderInformationGetResponse
	if err := clt.Do(req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Response.OrderInfo, nil
}
