package oauth

import (
	"context"

	"github.com/bububa/openpdd/api/ddk"
	"github.com/bububa/openpdd/core"
)

// ResourceUrlGenRequest 生成多多进宝频道推广 API Request
type ResourceUrlGenRequest struct {
	ddk.ResourceUrlGenRequest
}

// GetType implement Request interface
func (r ResourceUrlGenRequest) GetType() string {
	return "pdd.ddk.oauth.resource.url.gen"
}

// ResourceUrlGen 生成多多进宝频道推广
func ResourceUrlGen(ctx context.Context, clt *core.SDKClient, req *ResourceUrlGenRequest, accessToken string) (*ddk.PromURL, error) {
	var resp ddk.ResourceUrlGenResponse
	if err := clt.Do(ctx, req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Response, nil
}
