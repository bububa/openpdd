package oauth

import (
	"github.com/bububa/openpdd/api/ddk"
	"github.com/bububa/openpdd/core"
)

// WeappQrcodeUrlGenRequest 多多客生成单品推广小程序二维码url API Request
type WeappQrcodeUrlGenRequest struct {
	ddk.WeappQrcodeUrlGenRequest
}

// GetType implement Request interface
func (r WeappQrcodeUrlGenRequest) GetType() string {
	return "pdd.ddk.oauth.weapp.qrcode.url.gen"
}

// WeappQrcodeUrlGen 多多客生成单品推广小程序二维码url
func WeappQrcodeUrlGen(clt *core.SDKClient, req *WeappQrcodeUrlGenRequest, accessToken string) (string, error) {
	var resp ddk.WeappQrcodeUrlGenResponse
	if err := clt.Do(req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Response.URL, nil
}
