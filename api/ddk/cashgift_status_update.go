package ddk

import (
	"context"

	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
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
	return "pdd.ddk.cashgift.status.update"
}

// CashgiftStatusUpdateResponse 多多礼金状态更新 API Response
type CashgiftStatusUpdateResponse struct {
	model.CommonResponse
	Response struct {
		// ID 多多礼金ID
		ID uint64 `json:"cash_gift_id,omitempty" xml:"cash_gift_id,omitempty"`
	} `json:"update_cashgift_response" xml:"update_cashgift_response"`
}

// CashgiftStatusUpdate 多多礼金状态更新
func CashgiftStatusUpdate(ctx context.Context, clt *core.SDKClient, req *CashgiftStatusUpdateRequest) (uint64, error) {
	var resp CashgiftStatusUpdateResponse
	if err := clt.Do(ctx, req, &resp, ""); err != nil {
		return 0, err
	}
	return resp.Response.ID, nil
}
