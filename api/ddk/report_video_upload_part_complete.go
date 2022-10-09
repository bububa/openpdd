package ddk

import (
	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// ReportVideoUploadPartCompleteRequest 多多客信息流投放备案视频上传分片完成接口 API Request
type ReportVideoUploadPartCompleteRequest struct {
	// UploadSign 标记本次大文件上传的id（init阶段的返回值）
	UploadSign string `json:"upload_sign,omitempty"`
}

// GetType implement Request interface
func (r ReportVideoUploadPartCompleteRequest) GetType() string {
	return "pdd.ddk.report.video.upload.part.complete"
}

// Encode implement UploadRequest interface
func (r ReportVideoUploadPartCompleteRequest) Encode() []model.UploadField {
	return []model.UploadField{
		{
			Key:   "upload_sign",
			Value: r.UploadSign,
		},
	}
}

// ReportVideoUploadPartCompleteResponse 多多客信息流投放备案视频上传分片完成接口 API Response
type ReportVideoUploadPartCompleteResponse struct {
	model.CommonResponse
	Response struct {
		// URL 创建的视频资源对应的vid
		URL string `json:"url,omitempty"`
	} `json:"response"`
}

// ReportVidoeUploadPartComplete 多多客信息流投放备案视频上传分片完成接口
func ReportVideoUploadPartComplete(clt *core.SDKClient, uploadSign string) (string, error) {
	var resp ReportVideoUploadPartCompleteResponse
	if err := clt.Upload(&ReportVideoUploadPartCompleteRequest{
		UploadSign: uploadSign,
	}, &resp, ""); err != nil {
		return "", err
	}
	return resp.Response.URL, nil
}
