package ddk

import (
	"context"

	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// GoodsPidGenerateRequest 创建多多进宝推广位 API Request
type GoodsPidGenerateRequest struct {
	// Number 要生成的推广位数量，默认为10，范围为：1~100
	Number int `json:"number"`
	// PidNameList 推广位名称，例如["1","2"]
	PidNameList []string `json:"p_id_name_list,omitempty"`
	// MediaID 媒体id
	MediaID uint64 `json:"media_id,omitempty"`
}

// GetType implement Request interface
func (r GoodsPidGenerateRequest) GetType() string {
	return "pdd.ddk.goods.pid.generate"
}

// GoodsPidGenerateResponse 创建多多进宝推广位 API Response
type GoodsPidGenerateResponse struct {
	model.CommonResponse
	// Response response
	Response struct {
		// List 多多进宝推广位对象列表
		List []Pid `json:"p_id_list,omitempty"`
	} `json:"p_id_generate_response" xml:"p_id_generate_response"`
}

// GoodsPidGenerate 创建多多进宝推广位
func GoodsPidGenerate(ctx context.Context, clt *core.SDKClient, req *GoodsPidGenerateRequest) ([]Pid, error) {
	var resp GoodsPidGenerateResponse
	if err := clt.Do(ctx, req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Response.List, nil
}
