package pmc

import (
	"context"

	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// UserGetRequest 获取用户已开通消息 API Request
type UserGetRequest struct {
	// OwnerID 用户唯一id
	OwnerID string `json:"owner_id,omitempty"`
}

// GetType implement Request interface
func (r UserGetRequest) GetType() string {
	return "pdd.pmc.user.get"
}

// UserGetResponse 获取用户已开通消息 API Response
type UserGetResponse struct {
	model.CommonResponse
	Response struct {
		// User 开通的用户数据
		User *User `json:"pmc_user,omitempty"`
	} `json:"pmc_user_get_response"`
}

// User 开通的用户数据
type User struct {
	// Created 用户首次开通时间
	Created string `json:"created,omitempty"`
	// Modified 用户最后开通时间
	Modified string `json:"modified,omitempty"`
	// IsExpire 用户授权是否有效，0表示授权有效，1表示授权过期
	IsExpire int `json:"is_expire,omitempty"`
	// OwnerID 用户ID
	OwnerID string `json:"owner_id,omitempty"`
	// Topics 用户开通的消息类型列表。如果为空表示应用开通的所有类型
	Topics []string `json:"topics,omitempty"`
}

// UserGet 获取用户已开通消息
func UserGet(ctx context.Context, clt *core.SDKClient, ownerID string) (*User, error) {
	req := UserGetRequest{
		OwnerID: ownerID,
	}
	var resp UserGetResponse
	if err := clt.Do(ctx, &req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Response.User, nil
}
