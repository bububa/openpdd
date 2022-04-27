package pmc

import (
	"strings"

	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// UserPermitRequest 为已授权的用户开通消息服务 API Request
type UserPermitRequest struct {
	// Topics 消息主题列表，用半角逗号分隔。当用户订阅的topic是应用订阅的子集时才需要设置，不设置表示继承应用所订阅的所有topic，一般情况建议不要设置。
	Topics string `json:"topics,omitempty"`
}

// GetType implement Request interface
func (r UserPermitRequest) GetType() string {
	return "pdd.pmc.user.permit"
}

// UserPermitResponse 为已授权的用户开通消息服务 API Response
type UserPermitResponse struct {
	model.CommonResponse
	Response struct {
		// IsSuccess 是否成功
		IsSuccess bool `json:"is_success,omitempty"`
	} `json:"pmc_user_permit_response"`
}

// UserPermit 为已授权的用户开通消息服务
func UserPermit(clt *core.SDKClient, topics []string, accessToken string) (bool, error) {
	var req UserPermitRequest
	if len(topics) > 0 {
		req.Topics = strings.Join(topics, ",")
	}
	var resp UserPermitResponse
	if err := clt.Do(&req, &resp, accessToken); err != nil {
		return false, err
	}
	return resp.Response.IsSuccess, nil
}
