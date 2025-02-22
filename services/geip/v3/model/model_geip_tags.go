package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GeipTags 资源标签。
type GeipTags struct {

	// 标签键，最大长度128个unicode字符，格式为大小写字母，数字，中划线“-”，下划线“_”，中文。
	Key string `json:"key"`

	// tag的value列表
	Values []string `json:"values"`
}

func (o GeipTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GeipTags struct{}"
	}

	return strings.Join([]string{"GeipTags", string(data)}, " ")
}
