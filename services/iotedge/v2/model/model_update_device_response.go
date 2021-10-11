package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateDeviceResponse struct {
	// 设备配置，内容由产品的$config服务定义。

	Config         *interface{} `json:"config,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateDeviceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDeviceResponse struct{}"
	}

	return strings.Join([]string{"UpdateDeviceResponse", string(data)}, " ")
}
