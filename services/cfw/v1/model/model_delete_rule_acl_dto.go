package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteRuleAclDto struct {

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用查询防火墙实例接口获得，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。具体可参考APIExlorer和帮助中心FAQ。
	ObjectId string `json:"object_id"`

	// 规则列表
	RuleIds []string `json:"rule_ids"`
}

func (o DeleteRuleAclDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRuleAclDto struct{}"
	}

	return strings.Join([]string{"DeleteRuleAclDto", string(data)}, " ")
}
