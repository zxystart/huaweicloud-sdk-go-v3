package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListPublicBorderGroupsResponse struct {
	// 功能说明：公共池分组对象

	PublicBorderGroups *[]CommonPoolsWithBorderGroupDict `json:"public_border_groups,omitempty"`
	// 本次请求的编号

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPublicBorderGroupsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPublicBorderGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListPublicBorderGroupsResponse", string(data)}, " ")
}
