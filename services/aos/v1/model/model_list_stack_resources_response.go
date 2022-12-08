package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListStackResourcesResponse struct {

	// 资源栈输出
	StackResources *[]StackResource `json:"stack_resources,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListStackResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStackResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListStackResourcesResponse", string(data)}, " ")
}
