package ddk

import (
	"github.com/bububa/openpdd/core"
	"github.com/bububa/openpdd/model"
)

type PeroidType int

const (
	PeroidType_7DAY  PeroidType = 1
	PeroidType_MONTH PeroidType = 2
)

// StatisticDataQueryRequest 多多进宝数据统计查询接口 API Request
type StatisticDataQueryRequest struct {
	// Page 分页数，默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 每页结果数，默认值: 20
	PageSize int `json:"page_size,omitempty"`
	// PeriodType 周期类型: 1-每7天，2-自然月
	PeroidType PeroidType `json:"peroid_type,omitempty"`
	// StatisticType 数据类型: 1-增量补贴数据
	StatisticType int `json:"statistic_type,omitempty"`
	// Time 查询时间点，格式: "yyyy-MM-dd"。period_type为1时，查询时间点前7天的数据；period_type为2时，查询时间点所在自然月的数据。
	Time string `json:"time,omitempty"`
}

// GetType implement Request interface
func (r StatisticDataQueryRequest) GetType() string {
	return "pdd.ddk.statistics.data.query"
}

// StatisticDataQueryResponse 多多进宝数据统计查询接口 API Resposne
type StatisticDataQueryResponse struct {
	model.CommonResponse
	Response struct {
		List []StatisticData `json:"data_list,omitempty" xml:"data_list,omitempty"`
	} `json:"statistics_data_response" xml:"statistics_data_response"`
}

// StatisticData .
type StatisticData struct {
	// CustomParameters 自定义参数，为链接打上自定义标签；自定义参数最长限制64个字节；格式为： {"uid":"11111","sid":"22222"} ，其中 uid 用户唯一标识，可自行加密后传入，每个用户仅且对应一个标识，必填； sid 上下文信息标识，例如sessionId等，非必填。该json字符串中也可以加入其他自定义的key。（如果使用GET请求，请使用URLEncode处理参数）
	CustomParameters string `json:"custom_parameters,omitempty" xml:"custom_parameters,omitempty"`
	// EndTime 结束时间，格式: "yyyy-MM-dd"
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// StartTime 开始时间，格式: "yyyy-MM-dd"
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// OrderAmount GMV,单位为分
	OrderAmount int64 `json:"order_amount,omitempty" xml:"order_amount,omitempty"`
	// OrderNum 订单数
	OrderNum int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// Pid 推广位ID
	Pid string `json:"p_id,omitempty" xml:"p_id,omitempty"`
}

// StatisticDataQuery 多多进宝数据统计查询接口
func StatisticDataQuery(clt *core.SDKClient, req *StatisticDataQueryRequest) ([]StatisticData, error) {
	var resp StatisticDataQueryResponse
	if err := clt.Do(req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Response.List, nil
}
