package ddk

import (
	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

// MemberAuthorityQueryRequest 查询是否绑定备案 API Request
type MemberAuthorityQueryRequest struct {
	// CustomParameters 自定义参数，为链接打上自定义标签；自定义参数最长限制64个字节；格式为： {"uid":"11111","sid":"22222"} ，其中 uid 用户唯一标识，可自行加密后传入，每个用户仅且对应一个标识，必填； sid 上下文信息标识，例如sessionId等，非必填。该json字符串中也可以加入其他自定义的key。（如果使用GET请求，请使用URLEncode处理参数）
	CustomParameters interface{} `json:"custom_parameters,omitempty"`
	// Pid 推广位id
	Pid string `json:"pid,omitempty"`
}

// GetType implement Request interface
func (r MemberAuthorityQueryRequest) GetType() string {
	return "pdd.ddk.member.authority.query"
}

// MemberAuthorityQueryResponse 查询是否绑定备案 API Respose
type MemberAuthorityQueryResponse struct {
	model.CommonResponse
	// Response response
	Response struct {
		// Bind 1-已绑定；0-未绑定
		Bind int `json:"bind,omitempty" xml:"bind,omitempty"`
	} `json:"authority_query_response" xml:"authority_query_response"`
}

// MemberAuthorityQuery 查询是否绑定备案
func MemberAuthorityQuery(clt *core.SDKClient, req *MemberAuthorityQueryRequest, accessToken string) (bool, error) {
	var resp MemberAuthorityQueryResponse
	if err := clt.Do(req, &resp, accessToken); err != nil {
		return false, err
	}
	return resp.Response.Bind == 1, nil
}
