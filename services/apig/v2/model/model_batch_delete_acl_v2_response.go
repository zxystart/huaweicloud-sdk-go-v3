package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchDeleteAclV2Response struct {
	// 删除成功的ACL策略数量

	SuccessCount *int32 `json:"success_count,omitempty"`
	// 删除失败的ACL策略及错误信息

	Failure        *[]AclBatchResultFailureResp `json:"failure,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o BatchDeleteAclV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteAclV2Response struct{}"
	}

	return strings.Join([]string{"BatchDeleteAclV2Response", string(data)}, " ")
}
