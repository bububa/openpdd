package oauth

import (
	"github.com/bububa/openpdd/api/ddk"
	"github.com/bububa/openpdd/core"
)

// CashgiftStatusUpdateRequest 多多礼金状态更新 API Request
type CashgiftStatusUpdateRequest struct {
	ddk.CashgiftStatusUpdateRequest
}

// GetType implement Request interface
func (r CashgiftStatusUpdateRequest) GetType() string {
	return "pdd.ddk.oauth.cashgift.status.update"
}

// CashgiftStatusUpdate 多多礼金状态更新
func CashgiftStatusUpdate(clt *core.SDKClient, req *CashgiftStatusUpdateRequest, accessToken string) (uint64, error) {
	var resp ddk.CashgiftStatusUpdateResponse
	if err := clt.Do(req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Response.ID, nil
}
