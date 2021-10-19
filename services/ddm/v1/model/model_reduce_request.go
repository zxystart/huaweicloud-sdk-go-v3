package model

import (
	"encoding/json"

	"strings"
)

// This is a auto request Object
type ReduceRequest struct {
	// 需要缩容的节点数量，最大值为实例节点数减1。

	NodeNumber int32 `json:"node_number"`
}

func (o ReduceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ReduceRequest struct{}"
	}

	return strings.Join([]string{"ReduceRequest", string(data)}, " ")
}
