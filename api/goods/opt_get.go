package goods

import (
	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// OptGetRequest 查询商品标签列表 API Request
type OptGetRequest struct {
	// ParentCatID 值=0时为顶点cat_id,通过树顶级节点获取cat树
	ParentCatID uint64 `json:"parent_cat_id"`
}

// GetType implement Request interface
func (r OptGetRequest) GetType() string {
	return "pdd.goods.opt.get"
}

// OptGetResponse 查询商品标签列表 API Response
type OptGetResponse struct {
	model.CommonResponse
	Response struct {
		List []Opt `json:"goods_opt_list,omitempty"`
	} `json:"goods_opt_get_response"`
}

// OptGet 查询商品标签列表
func OptGet(clt *core.SDKClient, parentID uint64) ([]Opt, error) {
	req := OptGetRequest{
		ParentCatID: parentID,
	}
	var resp OptGetResponse
	if err := clt.Do(req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Response.List, nil
}
