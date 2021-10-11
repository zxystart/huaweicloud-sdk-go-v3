package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteDeviceRequest struct {
	// 边缘节点ID

	EdgeNodeId string `json:"edge_node_id"`
	// 设备ID

	DeviceId string `json:"device_id"`
}

func (o DeleteDeviceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteDeviceRequest struct{}"
	}

	return strings.Join([]string{"DeleteDeviceRequest", string(data)}, " ")
}
