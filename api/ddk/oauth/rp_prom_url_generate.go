package oauth

import (
	"context"

	"github.com/bububa/openpdd/api/ddk"
	"github.com/bububa/openpdd/core"
)

// RpPromUrlGenerateRequest 生成营销工具推广链接 API Request
type RpPromUrlGenerateRequest struct {
	ddk.RpPromUrlGenerateRequest
}

// GetType implement Request interface
func (r RpPromUrlGenerateRequest) GetType() string {
	return "pdd.ddk.oauth.rp.prom.url.generate"
}

// RpPromUrlGenerate 生成营销工具推广链接
func RpPromUrlGenerate(ctx context.Context, clt *core.SDKClient, req *RpPromUrlGenerateRequest, accessToken string) (*ddk.RpPromUrlGenerateResult, error) {
	var resp ddk.RpPromUrlGenerateResponse
	if err := clt.Do(ctx, req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Response, nil
}
