package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateInstallCmdRequest struct {
	// 边缘节点ID

	EdgeNodeId string `json:"edge_node_id"`
	// 节点架构

	Arch string `json:"arch"`
}

func (o CreateInstallCmdRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateInstallCmdRequest struct{}"
	}

	return strings.Join([]string{"CreateInstallCmdRequest", string(data)}, " ")
}
