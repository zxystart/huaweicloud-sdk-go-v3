package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListProjectCofigsV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0

	Offset *int64 `json:"offset,omitempty"`
	// 每页显示的条目数量

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListProjectCofigsV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListProjectCofigsV2Request struct{}"
	}

	return strings.Join([]string{"ListProjectCofigsV2Request", string(data)}, " ")
}
