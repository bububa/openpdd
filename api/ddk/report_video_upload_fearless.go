package ddk

import (
	"bufio"
	"bytes"
	"context"
	"sync"
	"sync/atomic"

	"github.com/bububa/openpdd/core"
)

// ReportVideoUploadFearless 多多客信息流投放备案视频上传接口
// 多多客信息流投放备案视频上传,上传视频大小有限制,单个文件超过20M需要走分片上传
// 根据上传文件大小自动选择API接口
func ReportVideoUploadFearless(ctx context.Context, clt *core.SDKClient, req *ReportVideoUploadRequest, size int64, parallel int) (string, error) {
	var maxSize int64 = 20 * 1 << 20
	var chunkSize int64 = 5 * 1 << 20
	if size <= maxSize {
		return ReportVideoUpload(ctx, clt, req)
	}
	uploadSign, err := ReportVideoUploadPartInit(ctx, clt, "video/mp4")
	if err != nil {
		return "", err
	}
	if parallel <= 0 {
		parallel = 1
	}
	var (
		fr        = bufio.NewReader(req.File)
		b         = make([]byte, chunkSize)
		wg        sync.WaitGroup
		guard     = make(chan struct{}, parallel)
		uploadErr = &atomic.Value{}
		partNum   int
	)
	for {
		n, err := fr.Read(b)
		if err != nil {
			break
		}
		partNum++
		buf := bytes.NewReader(b[0:n])
		partReq := ReportVideoUploadPartRequest{
			PartFile:   buf,
			PartNum:    partNum,
			UploadSign: uploadSign,
		}
		guard <- struct{}{}
		wg.Add(1)
		go func(req *ReportVideoUploadPartRequest) {
			defer wg.Done()
			if _, err = ReportVideoUploadPart(ctx, clt, &partReq); err != nil {
				uploadErr.Store(err)
			}
			<-guard
		}(&partReq)
	}
	wg.Wait()
	if err := uploadErr.Load(); err != nil {
		return "", err.(error)
	}
	return ReportVideoUploadPartComplete(ctx, clt, uploadSign)
}
