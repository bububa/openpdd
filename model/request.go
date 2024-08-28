package model

import "io"

// RequestDataType 请求返回的数据格式，可选参数为XML或JSON，默认为JSON
type RequestDataType string

const (
	// RequestDataType_JSON 返回JSON格式
	RequestDataType_JSON RequestDataType = "JSON"
	// RequestDataType_XML 返回X M Lg格式
	RequestDataType_XML RequestDataType = "XML"
)

// Request 请求
type Request interface {
	GetType() string
}

// CommonRequest 请求公共参数
type CommonRequest struct {
	// Type API接口名，形如：pdd.*
	Type string `json:"type,omitempty"`
	// ClientID 已创建成功的应用标志client_id，可在应用详情和中查看
	ClientID string `json:"client_id,omitempty"`
	// Sign API入参参数签名，签名值根据如下算法给出计算过程
	Sign string `json:"sign,omitempty"`
	// AccessToken 用户授权令牌access_token，通过pdd.pop.auth.token.create获取
	AccessToken string `json:"access_token,omitempty"`
	// Version API版本，默认为V1，无要求不传此参数
	Version string `json:"version,omitempty"`
	// DataType 请求返回的数据格式，可选参数为XML或JSON，默认为JSON
	DataType RequestDataType `json:"data_type,omitempty"`
	// Timestamp 时间戳，格式为UNIX时间（秒）
	Timestamp int64 `json:"timestamp,omitempty"`
}

// GetType implement Request interface
func (r CommonRequest) GetType() string {
	return r.Type
}

// UploadField multipart/form-data post request field struct
type UploadField struct {
	// Reader upload file reader
	Reader io.Reader
	// Key field key
	Key string
	// Value field value
	Value string
}

// UploadRequest multipart/form-data reqeust interface
type UploadRequest interface {
	// Encode encode request to UploadFields
	Encode() []UploadField
	GetType() string
}
