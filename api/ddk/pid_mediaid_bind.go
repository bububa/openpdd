package ddk

import (
	"context"

	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// PidMediaIDBindRequest 批量绑定推广位的媒体id API Request
type PidMediaIDBindRequest struct {
	// MediaID 媒体id
	MediaID uint64 `json:"media_id"`
	// PidList 推广位列表，例如：["60005_612"]，最多支持同时传入1000个
	PidList []string `json:"pid_list"`
}

// GetType implement Request interface
func (r PidMediaIDBindRequest) GetType() string {
	return "pdd.ddk.pid.mediaid.bind"
}

// PidMediaIDBindResponse 批量绑定推广位的媒体id API Response
type PidMediaIDBindResponse struct {
	model.CommonResponse
	// Response response
	Response struct {
		// Result 绑定结果
		Result *PidMediaIDBindResult `json:"result,omitempty" xml:"result,omitempty"`
	} `json:"p_id_bind_response" xml:"p_id_bind_response"`
}

// PidMediaIDBindResult 绑定结果
type PidMediaIDBindResult struct {
	// Msg 绑定结果文本提示
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// Result 绑定结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// PidMediaIDBind 批量绑定推广位的媒体id
func PidMediaIDBind(ctx context.Context, clt *core.SDKClient, mediaID uint64, pidList []string) (*PidMediaIDBindResult, error) {
	var (
		req = &PidMediaIDBindRequest{
			MediaID: mediaID,
			PidList: pidList,
		}
		resp PidMediaIDBindResponse
	)
	if err := clt.Do(ctx, req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Response.Result, nil
}
