package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListAclPolicyBindedToApiV2Response struct {
	// 本次返回的列表长度

	Size int32 `json:"size"`
	// 满足条件的记录数

	Total int64 `json:"total"`
	// 本次查询返回的ACL列表

	Acls           *[]ApiBindAclPageAclsResp `json:"acls,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListAclPolicyBindedToApiV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAclPolicyBindedToApiV2Response struct{}"
	}

	return strings.Join([]string{"ListAclPolicyBindedToApiV2Response", string(data)}, " ")
}
