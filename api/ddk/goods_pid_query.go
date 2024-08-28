package ddk

import (
	"context"

	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// GoodsPidQueryRequest 查询已经生成的推广位信息 API Request
type GoodsPidQueryRequest struct {
	// Page 返回的页数
	Page int `json:"page,omitempty"`
	// PageSize 返回的每页推广位数量
	PageSize int `json:"page_size,omitempty"`
	// PidList 推广位列表，例如：["60005_612"]
	PidList []string `json:"pid_list,omitempty"`
	// Status 推广位状态：0-正常，1-封禁
	Status *int `json:"status,omitempty"`
}

// GetType implement Request interface
func (r GoodsPidQueryRequest) GetType() string {
	return "pdd.ddk.goods.pid.query"
}

// GoodsPidQueryResponse 查询已经生成的推广位信息 API Response
type GoodsPidQueryResponse struct {
	model.CommonResponse
	Response struct {
		// List 多多进宝推广位对象列表
		List []Pid `json:"p_id_list,omitempty"`
		// TotalCount 返回推广位总数
		TotalCount int `json:"total_count,omitempty"`
	} `json:"p_id_query_response"`
}

// GoodsPidQuery 查询已经生成的推广位信息
func GoodsPidQuery(ctx context.Context, clt *core.SDKClient, req *GoodsPidQueryRequest) (int, []Pid, error) {
	var resp GoodsPidQueryResponse
	if err := clt.Do(ctx, req, &resp, ""); err != nil {
		return 0, nil, err
	}
	return resp.Response.TotalCount, resp.Response.List, nil
}
