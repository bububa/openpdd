package oauth

import (
	"github.com/bububa/openpdd/api/ddk"
	"github.com/bububa/openpdd/core"
)

// CashgiftStatusUpdateRequest 多多礼金状态更新 API Request
type CashgiftStatusUpdateRequest struct {
	// ID 多多礼金ID
	ID uint64 `json:"cash_gift_id"`
	// UpdateType 礼金更新类型：0-停止礼金推广，1-恢复礼金推广
	UpdateType int `json:"update_type"`
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
