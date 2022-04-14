package oauth

import (
	"github.com/bububa/openpdd/api/ddk"
	"github.com/bububa/openpdd/core"
)

// GoodsPidGenerateRequest 创建多多进宝推广位 API Request
type GoodsPidGenerateRequest struct {
	ddk.GoodsPidGenerateRequest
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
