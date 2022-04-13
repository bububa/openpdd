package oauth

import (
	"github.com/bububa/openpdd/api/ddk"
	"github.com/bububa/openpdd/core"
)

// GoodsPidQueryRequest 查询已经生成的推广位信息 API Request
type GoodsPidQueryRequest struct {
	// Page 返回的页数
	Page int `json:"page,omitempty"`
	// PageSize 返回的每页推广位数量
	PageSize int `json:"page_size,omitempty"`
	// PidList 推广位列表，例如：["60005_612"]
	PidList []string `json:"pid_list,omitempty"`
}

// GetType implement Request interface
func (r GoodsPidQueryRequest) GetType() string {
	return "pdd.ddk.oauth.goods.pid.query"
}

// GoodsPidQuery 查询已经生成的推广位信息
func GoodsPidQuery(clt *core.SDKClient, req *GoodsPidQueryRequest, accessToken string) (int, []ddk.Pid, error) {
	var resp ddk.GoodsPidQueryResponse
	if err := clt.Do(req, &resp, accessToken); err != nil {
		return 0, nil, err
	}
	return resp.Response.TotalCount, resp.Response.List, nil
}
