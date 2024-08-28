package oauth

import (
	"context"

	"github.com/bububa/openpdd/api/ddk"
	"github.com/bububa/openpdd/core"
)

// GoodsPidQueryRequest 查询已经生成的推广位信息 API Request
type GoodsPidQueryRequest struct {
	ddk.GoodsPidQueryRequest
}

// GetType implement Request interface
func (r GoodsPidQueryRequest) GetType() string {
	return "pdd.ddk.oauth.goods.pid.query"
}

// GoodsPidQuery 查询已经生成的推广位信息
func GoodsPidQuery(ctx context.Context, clt *core.SDKClient, req *GoodsPidQueryRequest, accessToken string) (int, []ddk.Pid, error) {
	var resp ddk.GoodsPidQueryResponse
	if err := clt.Do(ctx, req, &resp, accessToken); err != nil {
		return 0, nil, err
	}
	return resp.Response.TotalCount, resp.Response.List, nil
}
