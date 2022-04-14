package oauth

import (
	"github.com/bububa/openpdd/api/ddk"
	"github.com/bububa/openpdd/core"
)

// OrderDetailGetRequest 查询订单详情 API Request
type OrderDetailGetRequest struct {
	ddk.OrderDetailGetRequest
}

// GetType implement Request interface
func (r OrderDetailGetRequest) GetType() string {
	return "pdd.ddk.oauth.order.detail.get"
}

// OrderDetailGet 查询订单详情
func OrderDetailGet(clt *core.SDKClient, req *OrderDetailGetRequest, accessToken string) (*ddk.Order, error) {
	var resp ddk.OrderDetailGetResponse
	if err := clt.Do(req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Response, nil
}
