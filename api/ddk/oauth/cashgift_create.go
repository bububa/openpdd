package oauth

import (
	"github.com/bububa/openpdd/api/ddk"
	"github.com/bububa/openpdd/core"
)

// CashgiftCreateRequest 创建多多礼金 API Request
type CashgiftCreateRequest struct {
	ddk.CashgiftCreateRequest
}

// GetType implement Request interface
func (r CashgiftCreateRequest) GetType() string {
	return "pdd.ddk.oauth.cashgift.create"
}

// CashgiftCreate 创建多多礼金
func CashgiftCreate(clt *core.SDKClient, req *CashgiftCreateRequest, accessToken string) (*ddk.CashgiftCreateResult, error) {
	var resp ddk.CashgiftCreateResponse
	if err := clt.Do(req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Response, nil
}
