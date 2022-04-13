package auth

import (
	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// TokenCreateRequest 获取Access Token API Request
type TokenCreateRequest struct {
	// Code 授权code，grantType==authorization_code 时需要
	Code string `json:"code,omitempty"`
}

// GetType implement Request interface
func (r TokenCreateRequest) GetType() string {
	return "pdd.pop.auth.token.create"
}

// TokenCreateResponse 获取Access Token API Response
type TokenCreateResponse struct {
	model.CommonResponse
	// Response response
	Response *Token `json:"pop_auth_token_create_response,omitempty" xml:"pop_auth_token_create_response,omitempty"`
}

// TokenCreate 获取Access Token
func TokenCreate(clt *core.SDKClient, code string) (*Token, error) {
	req := &TokenCreateRequest{
		Code: code,
	}
	var resp TokenCreateResponse
	if err := clt.Do(req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Response, nil
}
