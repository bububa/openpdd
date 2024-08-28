package goods

import (
	"context"

	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// CatsGetRequest 商品标准类目接口 API Request
type CatsGetRequest struct {
	// ParentCatID 值=0时为顶点cat_id,通过树顶级节点获取cat树
	ParentCatID uint64 `json:"parent_cat_id"`
}

// GetType implement Request interface
func (r CatsGetRequest) GetType() string {
	return "pdd.goods.cats.get"
}

// CatsGetResponse 商品标准类目接口 API Response
type CatsGetResponse struct {
	model.CommonResponse
	Response struct {
		List []Cat `json:"goods_cats_list,omitempty"`
	} `json:"goods_cats_get_response"`
}

// CatsGet 商品标准类目接口
func CatsGet(ctx context.Context, clt *core.SDKClient, parentID uint64) ([]Cat, error) {
	req := CatsGetRequest{
		ParentCatID: parentID,
	}
	var resp CatsGetResponse
	if err := clt.Do(ctx, req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Response.List, nil
}
