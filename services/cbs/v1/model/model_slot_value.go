package model

import (
	"encoding/json"

	"strings"
)

//
type SlotValue struct {
	// 词。

	Word string `json:"word"`
	// 归一化后的标准词。

	NormWord string `json:"norm_word"`
	// 词的起始位置。

	BeginPosition int32 `json:"begin_position"`
	// 词的结束位置。

	EndPosition int32 `json:"end_position"`
}

func (o SlotValue) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SlotValue struct{}"
	}

	return strings.Join([]string{"SlotValue", string(data)}, " ")
}
