package auth

import (
	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// TokenRefreshRequest 刷新Access Token API Request
type TokenRefreshRequest struct {
	// RefreshToken grantType==refresh_token 时需要
	RefreshToken string `json:"refresh_token,omitempty"`
}

// GetType implement Request interface
func (r TokenRefreshRequest) GetType() string {
	return "pdd.pop.auth.token.refresh"
}

// TokenRefreshResponse 刷新Access Token API Response
type TokenRefreshResponse struct {
	model.CommonResponse
	Data *Token `json:"pop_auth_token_refresh_response,omitempty" xml:"pop_auth_token_refresh_response,omitempty"`
}

// TokenRefresh 刷新Access Token
func TokenRefresh(clt *core.SDKClient, refreshToken string) (*Token, error) {
	req := &TokenRefreshRequest{
		RefreshToken: refreshToken,
	}
	var resp TokenRefreshResponse
	if err := clt.Do(req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
