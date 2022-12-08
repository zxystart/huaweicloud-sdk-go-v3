package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListStackOutputsResponse struct {

	// 资源栈输出
	Outputs        *[]StackOutput `json:"outputs,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListStackOutputsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStackOutputsResponse struct{}"
	}

	return strings.Join([]string{"ListStackOutputsResponse", string(data)}, " ")
}
