package ddk

import (
	"io"

	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// ReportImgUploadRequest 多多客信息流投放备案图片上传接口 API Request
type ReportImgUploadRequest struct {
	// File 多多视频图片文件流
	File io.Reader `json:"file,omitempty"`
	// Filename
	Filename string `json:"filename,omitempty"`
}

// GetType implement Request interface
func (r ReportImgUploadRequest) GetType() string {
	return "pdd.ddk.report.img.upload"
}

// Encode implement UploadRequest interface
func (r ReportImgUploadRequest) Encode() []model.UploadField {
	return []model.UploadField{
		{
			Key:    "file",
			Value:  r.Filename,
			Reader: r.File,
		},
	}
}

// ReportImgUploadResponse 多多客信息流投放备案图片上传接口 API Response
type ReportImgUploadResponse struct {
	model.CommonResponse
	Response struct {
		// URL
		URL string `json:"url,omitempty"`
	} `json:"response"`
}

// ReportImgUpload 多多进宝信息流渠道备案授权素材上传接口
func ReportImgUpload(clt *core.SDKClient, req *ReportImgUploadRequest) (string, error) {
	var resp ReportImgUploadResponse
	if err := clt.Upload(req, &resp, ""); err != nil {
		return "", err
	}
	return resp.Response.URL, nil
}
