package pmc

import (
	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// UserCancelRequest 取消用户的消息服务 API Request
type UserCancelRequest struct {
	// OwnerID 用户唯一id
	OwnerID string `json:"owner_id,omitempty"`
}

// GetType implement Request interface
func (r UserCancelRequest) GetType() string {
	return "pdd.pmc.user.cancel"
}

// UserCancelResponse 取消用户的消息服务 API Response
type UserCancelResponse struct {
	model.CommonResponse
	Response struct {
		// IsSuccess 是否成功
		IsSuccess bool `json:"is_success,omitempty"`
	} `json:"tmc_user_cancel_response"`
}

// UserCancel 取消用户的消息服务
func UserCancel(clt *core.SDKClient, ownerID string) (bool, error) {
	req := UserCancelRequest{
		OwnerID: ownerID,
	}
	var resp UserCancelResponse
	if err := clt.Do(&req, &resp, ""); err != nil {
		return false, err
	}
	return resp.Response.IsSuccess, nil
}
