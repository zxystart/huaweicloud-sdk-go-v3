package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateRoutesResponse struct {
	// 路由列表

	Routes *[]RouterRespDto `json:"routes,omitempty"`
	// 最后一次修改时间

	UpdateTime     *string `json:"update_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateRoutesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateRoutesResponse struct{}"
	}

	return strings.Join([]string{"UpdateRoutesResponse", string(data)}, " ")
}
