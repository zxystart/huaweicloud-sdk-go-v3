package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListLatelyGroupStatisticsV2Response struct {
	// 响应码

	Code *string `json:"code,omitempty"`
	// 返回消息

	Msg *string `json:"msg,omitempty"`
	// 开始时间的UTC的时间戳

	StartTime *int32 `json:"start_time,omitempty"`
	// 截止时间的UTC的时间戳

	EndTime *int32 `json:"end_time,omitempty"`
	// 统计指标的数据结构结构体

	List           *[]StatisticsGroup `json:"list,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListLatelyGroupStatisticsV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListLatelyGroupStatisticsV2Response struct{}"
	}

	return strings.Join([]string{"ListLatelyGroupStatisticsV2Response", string(data)}, " ")
}
