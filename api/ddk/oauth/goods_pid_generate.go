package oauth

import (
	"github.com/bububa/openpdd/api/ddk"
	"github.com/bububa/openpdd/core"
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
	return "pdd.ddk.oauth.goods.pid.generate"
}

// GoodsPidGenerate 创建多多进宝推广位
func GoodsPidGenerate(clt *core.SDKClient, req *GoodsPidGenerateRequest, accessToken string) ([]ddk.Pid, error) {
	var resp ddk.GoodsPidGenerateResponse
	if err := clt.Do(req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Response.List, nil
}
