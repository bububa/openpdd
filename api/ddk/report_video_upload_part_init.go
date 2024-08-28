package ddk

import (
	"context"

	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// ReportVideoUploadPartInitRequest 多多客信息流投放备案视频上传分片初始化接口 API Request
type ReportVideoUploadPartInitRequest struct {
	// ContentType 文件对应的contentType，且必须为视频类型
	ContentType string `json:"content_type,omitempty"`
}

// GetType implement Request interface
func (r ReportVideoUploadPartInitRequest) GetType() string {
	return "pdd.ddk.report.video.upload.part.init"
}

// Encode implement UploadRequest interface
func (r ReportVideoUploadPartInitRequest) Encode() []model.UploadField {
	return []model.UploadField{
		{
			Key:   "content_type",
			Value: r.ContentType,
		},
	}
}

// ReportVideoUploadPartInitResponse 多多客信息流投放备案视频上传分片初始化接口 API Response
type ReportVideoUploadPartInitResponse struct {
	model.CommonResponse
	Response struct {
		UploadSign string `json:"upload_sign,omitempty"`
	} `json:"response"`
}

// ReportVidoeUploadPartInit 多多客信息流投放备案视频上传分片初始化接口
func ReportVideoUploadPartInit(ctx context.Context, clt *core.SDKClient, contentType string) (string, error) {
	var resp ReportVideoUploadPartInitResponse
	if err := clt.Upload(ctx, &ReportVideoUploadPartInitRequest{
		ContentType: contentType,
	}, &resp, ""); err != nil {
		return "", err
	}
	return resp.Response.UploadSign, nil
}
