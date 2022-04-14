package oauth

import (
	"github.com/bububa/openpdd/api/ddk"
	"github.com/bububa/openpdd/core"
)

// CmsPromUrlGenerateRequest 生成商城-频道推广链接 API Request
type CmsPromUrlGenerateRequest struct {
	ddk.CmsPromUrlGenerateRequest
}

// GetType implement Request interface
func (r CmsPromUrlGenerateRequest) GetType() string {
	return "pdd.ddk.oauth.cms.prom.url.generate"
}

// CmsPromUrlGenerate 生成商城-频道推广链接
func CmsPromUrlGenerate(clt *core.SDKClient, req *CmsPromUrlGenerateRequest, accessToken string) (int, []ddk.PromURL, error) {
	var resp ddk.CmsPromUrlGenerateResponse
	if err := clt.Do(req, &resp, accessToken); err != nil {
		return 0, nil, err
	}
	return resp.Response.Total, resp.Response.URLList, nil
}
