package oauth

import (
	"github.com/bububa/openpdd/api/ddk"
	"github.com/bububa/openpdd/core"
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
	return "pdd.ddk.oauth.pid.mediaid.bind"
}

// PidMediaIDBind 批量绑定推广位的媒体id
func PidMediaIDBind(clt *core.SDKClient, mediaID uint64, pidList []string, accessToken string) (*ddk.PidMediaIDBindResult, error) {
	var (
		req = &PidMediaIDBindRequest{
			MediaID: mediaID,
			PidList: pidList,
		}
		resp ddk.PidMediaIDBindResponse
	)
	if err := clt.Do(req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Response.Result, nil
}
