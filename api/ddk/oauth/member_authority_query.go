package oauth

import (
	"github.com/bububa/openpdd/api/ddk"
	"github.com/bububa/openpdd/core"
)

// MemberAuthorityQueryRequest 查询是否绑定备案 API Request
type MemberAuthorityQueryRequest struct {
	ddk.MemberAuthorityQueryRequest
}

// GetType implement Request interface
func (r MemberAuthorityQueryRequest) GetType() string {
	return "pdd.ddk.oauth.member.authority.query"
}

// MemberAuthorityQuery 查询是否绑定备案
func MemberAuthorityQuery(clt *core.SDKClient, req *MemberAuthorityQueryRequest, accessToken string) (bool, error) {
	var resp ddk.MemberAuthorityQueryResponse
	if err := clt.Do(req, &resp, accessToken); err != nil {
		return false, err
	}
	return resp.Response.Bind == 1, nil
}
