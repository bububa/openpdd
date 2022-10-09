package ddk

import (
	"bufio"
	"bytes"

	"github.com/bububa/openpdd/core"
)

// ReportVideoUploadFearless 多多客信息流投放备案视频上传接口
// 多多客信息流投放备案视频上传,上传视频大小有限制,单个文件超过20M需要走分片上传
// 根据上传文件大小自动选择API接口
func ReportVideoUploadFearless(clt *core.SDKClient, req *ReportVideoUploadRequest, size int64) (string, error) {
	var chunkSize int64 = 20 * 1 << 20
	if size <= chunkSize {
		return ReportVideoUpload(clt, req)
	}
	uploadSign, err := ReportVideoUploadPartInit(clt, "video/mp4")
	if err != nil {
		return "", err
	}
	var (
		fr      = bufio.NewReader(req.File)
		b       = make([]byte, 0, chunkSize)
		partNum int
	)
	for {
		n, err := fr.Read(b)
		if err != nil {
			break
		}
		buf := bytes.NewReader(b[0:n])
		partReq := ReportVideoUploadPartRequest{
			PartFile:   buf,
			PartNum:    partNum,
			UploadSign: uploadSign,
		}
		partNum, err = ReportVideoUploadPart(clt, &partReq)
		if err != nil {
			return "", err
		}
	}
	return ReportVideoUploadPartComplete(clt, uploadSign)
}
