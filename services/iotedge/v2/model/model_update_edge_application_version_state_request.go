package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateEdgeApplicationVersionStateRequest struct {
	// 应用ID，应用唯一标识。

	EdgeAppId string `json:"edge_app_id"`
	// 应用版本,应用内版本唯一。

	Version string `json:"version"`

	Body *UpdateEdgeAppVersionStateDto `json:"body,omitempty"`
}

func (o UpdateEdgeApplicationVersionStateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateEdgeApplicationVersionStateRequest struct{}"
	}

	return strings.Join([]string{"UpdateEdgeApplicationVersionStateRequest", string(data)}, " ")
}
