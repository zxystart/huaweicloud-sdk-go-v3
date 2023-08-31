package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportHostToEnvironmentResponse Response Object
type ImportHostToEnvironmentResponse struct {

	// 响应状态
	Status *string `json:"status,omitempty"`

	// 被导入的主机ids
	Result         *[]string `json:"result,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ImportHostToEnvironmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportHostToEnvironmentResponse struct{}"
	}

	return strings.Join([]string{"ImportHostToEnvironmentResponse", string(data)}, " ")
}
