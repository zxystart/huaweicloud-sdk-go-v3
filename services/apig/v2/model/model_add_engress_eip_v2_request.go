package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AddEngressEipV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`

	Body *OpenEngressEipReq `json:"body,omitempty"`
}

func (o AddEngressEipV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddEngressEipV2Request struct{}"
	}

	return strings.Join([]string{"AddEngressEipV2Request", string(data)}, " ")
}
