package ddk

// Pid 推广位
type Pid struct {
	// CreateTime 推广位创建时间
	CreateTime int64 `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// Name 推广位名称
	Name string `json:"pid_name,omitempty" xml:"pid_name,omitempty"`
	// ID 调用方推广位ID
	ID string `json:"p_id,omitempty" xml:"p_id,omitempty"`
	// MediaID 媒体id
	MediaID uint64 `json:"media_id,omitempty" xml:"media_id,omitempty"`
	// Status 推广位状态：0-正常，1-封禁
	Status int `json:"status,omitempty"`
}
