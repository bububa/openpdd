package ddk

import (
	"bufio"
	"bytes"
	"sync"
	"sync/atomic"

	"github.com/bububa/openpdd/core"
)

// ReportVideoUploadFearless 多多客信息流投放备案视频上传接口
// 多多客信息流投放备案视频上传,上传视频大小有限制,单个文件超过20M需要走分片上传
// 根据上传文件大小自动选择API接口
func ReportVideoUploadFearless(clt *core.SDKClient, req *ReportVideoUploadRequest, size int64, parallel int) (string, error) {
	var maxSize int64 = 20 * 1 << 20
	var chunkSize int64 = 5 * 1 << 20
	if size <= maxSize {
		return ReportVideoUpload(clt, req)
	}
	uploadSign, err := ReportVideoUploadPartInit(clt, "video/mp4")
	if err != nil {
		return "", err
	}
	var (
		fr      = bufio.NewReader(req.File)
		b       = make([]byte, chunkSize)
		partNum int
	)
	if parallel > 1 {
		ch := make(chan *ReportVideoUploadPartRequest, parallel)
		go func(ch chan<- *ReportVideoUploadPartRequest) {
			for {
				n, err := fr.Read(b)
				if err != nil {
					break
				}
				partNum++
				buf := bytes.NewReader(b[0:n])
				partReq := &ReportVideoUploadPartRequest{
					PartFile:   buf,
					PartNum:    partNum,
					UploadSign: uploadSign,
				}
				ch <- partReq
			}
			close(ch)
		}(ch)
		uploadErr := &atomic.Value{}
		var wg sync.WaitGroup
		for req := range ch {
			partReq := req
			wg.Add(1)
			go func(partReq *ReportVideoUploadPartRequest) {
				defer wg.Done()
				if _, err = ReportVideoUploadPart(clt, partReq); err != nil {
					uploadErr.Store(err)
				}

			}(partReq)
		}
		wg.Wait()
		if err := uploadErr.Load(); err != nil {
			return "", err.(error)
		}
	} else {
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
			if _, err = ReportVideoUploadPart(clt, &partReq); err != nil {
				return "", err
			}
		}
	}
	return ReportVideoUploadPartComplete(clt, uploadSign)
}
