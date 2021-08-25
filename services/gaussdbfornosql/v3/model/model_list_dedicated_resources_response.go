package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListDedicatedResourcesResponse struct {
	// 总记录数。

	TotalCount *int32 `json:"total_count,omitempty"`
	// 专属资源信息列表。

	Resources      *[]ListDedicatedResourceResult `json:"resources,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListDedicatedResourcesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListDedicatedResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListDedicatedResourcesResponse", string(data)}, " ")
}
