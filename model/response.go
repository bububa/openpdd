package model

import (
	"strconv"
	"strings"
)

// Response 请求返回
type Response interface {
	IsError() bool
	Error() error
}

// ErrorResponse Response Error
type ErrorResponse struct {
	BaseResponse
	ErrorCode int    `json:"error_code" xml:"error_code"`
	ErrorMsg  string `json:"error_msg" xml:"error_msg"`
	SubMsg    string `json:"sub_msg" xml:"sub_msg"`
	SubCode   string `json:"sub_code" xml:"sub_code"`
}

// Error implement Error interface
func (e ErrorResponse) Error() string {
	var builder strings.Builder
	builder.WriteString("CODE:")
	builder.WriteString(strconv.Itoa(e.ErrorCode))
	builder.WriteString(", MSG:")
	builder.WriteString(e.ErrorMsg)
	if e.SubCode != "" {
		builder.WriteString(", SUB_CODE:")
		builder.WriteString(e.SubCode)
	}
	if e.SubMsg != "" {
		builder.WriteString(", SUB_MSG:")
		builder.WriteString(e.SubMsg)
	}
	return builder.String()
}

// CommonResponse API Response
type CommonResponse struct {
	ErrorResponse *ErrorResponse `json:"error_response,omitempty" xml:"error_response,omitempty"`
}

// IsError 判断Response是否Error
func (r CommonResponse) IsError() bool {
	return r.ErrorResponse != nil && r.ErrorResponse.ErrorCode != 0
}

// Error returns response error
func (r CommonResponse) Error() error {
	if !r.IsError() {
		return nil
	}
	return r.ErrorResponse
}

// BaseResponse .
type BaseResponse struct {
	// RequestID .
	RequestID string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}
