package pmc

import (
	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// AccrueQueryRequest 消息队列积压数量查询 API Request
type AccrueQueryRequest struct{}

// GetType implement Request interface
func (r AccrueQueryRequest) GetType() string {
	return "pdd.pmc.accure.query"
}

// AccrueQueryResponse 消息队列积压数量查询 API Response
type AccrueQueryResponse struct {
	model.CommonResponse
	Response struct {
		// Number 消息积压数量
		Number int64 `json:"number,omitempty"`
	} `json:"pmc_user_get_response"`
}

// AccrueQuery 消息队列积压数量查询
func AccureQuery(clt *core.SDKClient) (int64, error) {
	var resp AccrueQueryResponse
	if err := clt.Do(&AccrueQueryRequest{}, &resp, ""); err != nil {
		return 0, err
	}
	return resp.Response.Number, nil
}
