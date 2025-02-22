package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagKey 标记的密钥标识符或名称。
type TagKey struct {
}

func (o TagKey) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagKey struct{}"
	}

	return strings.Join([]string{"TagKey", string(data)}, " ")
}
