package ddk

import (
	"context"
	"io"

	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// ReportVideoUploadRequest 多多客信息流投放备案视频上传接口 API Request
type ReportVideoUploadRequest struct {
	// File 多多视频图片文件流
	File io.Reader `json:"file,omitempty"`
	// Filename
	Filename string `json:"filename,omitempty"`
}

// GetType implement Request interface
func (r ReportVideoUploadRequest) GetType() string {
	return "pdd.ddk.report.video.upload"
}

// Encode implement UploadRequest interface
func (r ReportVideoUploadRequest) Encode() []model.UploadField {
	return []model.UploadField{
		{
			Key:    "file",
			Value:  r.Filename,
			Reader: r.File,
		},
	}
}

// ReportVideoUploadResponse 多多客信息流投放备案视频上传接口 API Response
type ReportVideoUploadResponse struct {
	model.CommonResponse
	Response struct {
		// URL
		URL string `json:"url,omitempty"`
	} `json:"response"`
}

// ReportVideoUpload 多多客信息流投放备案视频上传接口
// 多多客信息流投放备案视频上传,上传视频大小有限制,单个文件超过20M需要走分片上传
func ReportVideoUpload(ctx context.Context, clt *core.SDKClient, req *ReportVideoUploadRequest) (string, error) {
	var resp ReportVideoUploadResponse
	if err := clt.Upload(ctx, req, &resp, ""); err != nil {
		return "", err
	}
	return resp.Response.URL, nil
}
