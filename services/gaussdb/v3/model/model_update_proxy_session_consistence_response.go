package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateProxySessionConsistenceResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateProxySessionConsistenceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProxySessionConsistenceResponse struct{}"
	}

	return strings.Join([]string{"UpdateProxySessionConsistenceResponse", string(data)}, " ")
}
