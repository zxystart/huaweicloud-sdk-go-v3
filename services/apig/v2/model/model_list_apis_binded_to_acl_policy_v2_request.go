package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListApisBindedToAclPolicyV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// ACL编号

	AclId string `json:"acl_id"`
	// API编号

	ApiId *string `json:"api_id,omitempty"`
	// API名称

	ApiName *string `json:"api_name,omitempty"`
	// 环境编号

	EnvId *string `json:"env_id,omitempty"`
	// API分组编号

	GroupId *string `json:"group_id,omitempty"`
	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0

	Offset *int64 `json:"offset,omitempty"`
	// 每页显示的条目数量

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListApisBindedToAclPolicyV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListApisBindedToAclPolicyV2Request struct{}"
	}

	return strings.Join([]string{"ListApisBindedToAclPolicyV2Request", string(data)}, " ")
}
