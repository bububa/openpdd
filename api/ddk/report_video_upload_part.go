package ddk

import (
	"context"
	"io"
	"strconv"

	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// ReportVideoUploadPartRequest 多多客信息流投放备案视频上传分片上传接口 API Request
type ReportVideoUploadPartRequest struct {
	// PartFile 当前分片的文件流
	PartFile io.Reader `json:"part_file,omitempty"`
	// PartNum 当前分片编号名，从1开始
	PartNum int `json:"part_num,omitempty"`
	// UploadSign 标记本次大文件上传的id（init阶段的返回值）
	UploadSign string `json:"upload_sign,omitempty"`
}

// GetType implement Request interface
func (r ReportVideoUploadPartRequest) GetType() string {
	return "pdd.ddk.report.video.upload.part"
}

// Encode implement UploadRequest interface
func (r ReportVideoUploadPartRequest) Encode() []model.UploadField {
	partNum := r.PartNum
	if partNum == 0 {
		partNum = 1
	}
	return []model.UploadField{
		{
			Key:    "part_file",
			Reader: r.PartFile,
		},
		{
			Key:   "part_num",
			Value: strconv.Itoa(partNum),
		},
		{
			Key:   "upload_sign",
			Value: r.UploadSign,
		},
	}
}

// ReportVideoUploadPartResponse 多多客信息流投放备案视频上传分片上传接口 API Response
type ReportVideoUploadPartResponse struct {
	model.CommonResponse
	Response struct {
		// UploadPartNum 表示本次成功上传的part number
		UploadPartNum int `json:"uploaded_part_num,omitempty"`
	} `json:"response"`
}

// ReportVidoeUploadPart 多多客信息流投放备案视频上传分片上传接口
func ReportVideoUploadPart(ctx context.Context, clt *core.SDKClient, req *ReportVideoUploadPartRequest) (int, error) {
	var resp ReportVideoUploadPartResponse
	if err := clt.Upload(ctx, req, &resp, ""); err != nil {
		return 0, err
	}
	return resp.Response.UploadPartNum, nil
}
