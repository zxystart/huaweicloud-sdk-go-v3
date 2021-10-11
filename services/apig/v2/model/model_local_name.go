package model

import (
	"encoding/json"

	"strings"
)

// 可用区中英文名称。
type LocalName struct {
	// 可用区英文名称。

	EnUs *string `json:"en_us,omitempty"`
	// 可用区中文名称。

	ZhCn *string `json:"zh_cn,omitempty"`
}

func (o LocalName) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "LocalName struct{}"
	}

	return strings.Join([]string{"LocalName", string(data)}, " ")
}
