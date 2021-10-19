package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListServicePermissionsDetailsResponse struct {
	// permission列表。

	Permissions *[]Permissions `json:"permissions,omitempty"`
	// 满足查询条件的终端节点服务的白名单总条数，不受分页（即limit、offset参数）影响。

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListServicePermissionsDetailsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListServicePermissionsDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListServicePermissionsDetailsResponse", string(data)}, " ")
}
