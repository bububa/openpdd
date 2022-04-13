package auth

import (
	"github.com/bububa/openpdd/model"
)

// Token 调用令牌
type Token struct {
	model.BaseResponse
	OwnerName             string           `json:"owner_name,omitempty" xml:"owner_name,omitempty"`
	OwnerID               model.FlexUint64 `json:"owner_id,omitempty" xml:"owner_id,omitempty"`
	AccessToken           string           `json:"access_token,omitempty" xml:"access_token,omitempty"`
	RefreshToken          string           `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	ExpiresIn             int64            `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	ExpiresAt             int64            `json:"expires_at,omitempty" xml:"expires_at,omitempty"`
	RefreshTokenExpiresIn int64            `json:"refresh_token_expires_in,omitempty" xml:"refresh_token_expires_in,omitempty"`
	RefreshTokenExpiresAt int64            `json:"refresh_token_expires_at,omitempty" xml:"refresh_token_expires_at,omitempty"`
	W1ExpiresIn           int64            `json:"w1_expires_in,omitempty" xml:"w1_expires_in,omitempty"`
	W1ExpiresAt           int64            `json:"w1_expires_at,omitempty" xml:"w1_expires_at,omitempty"`
	W2ExpiresIn           int64            `json:"w2_expires_in,omitempty" xml:"w2_expires_in,omitempty"`
	W2ExpiresAt           int64            `json:"w2_expires_at,omitempty" xml:"w2_expires_at,omitempty"`
	R1ExpiresIn           int64            `json:"r1_expires_in,omitempty" xml:"r1_expires_in,omitempty"`
	R1ExpiresAt           int64            `json:"r1_expires_at,omitempty" xml:"r1_expires_at,omitempty"`
	R2ExpiresAt           int64            `json:"r2_expires_at,omitempty" xml:"r2_expires_at,omitempty"`
	R2ExpiresIn           int64            `json:"r2_expires_in,omitempty" xml:"r2_expires_in,omitempty"`
	Scope                 []string         `json:"scope,omitempty" xml:"scope,omitempty"`
}
