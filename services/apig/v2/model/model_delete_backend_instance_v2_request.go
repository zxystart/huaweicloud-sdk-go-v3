package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteBackendInstanceV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// VPC通道的编号

	VpcChannelId string `json:"vpc_channel_id"`
	// 后端实例对象的编号

	MemberId string `json:"member_id"`
}

func (o DeleteBackendInstanceV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteBackendInstanceV2Request struct{}"
	}

	return strings.Join([]string{"DeleteBackendInstanceV2Request", string(data)}, " ")
}
