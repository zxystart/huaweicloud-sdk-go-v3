package model

import (
	"encoding/json"

	"strings"
)

//
type CurrentSlot struct {
	// 槽位ID。

	SlotId string `json:"slot_id"`
	// 槽位名称。

	SlotName string `json:"slot_name"`
	// 槽位值。

	SlotValues []SlotValue `json:"slot_values"`
	// 槽位标识。

	SlotIdentification *string `json:"slot_identification,omitempty"`
}

func (o CurrentSlot) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CurrentSlot struct{}"
	}

	return strings.Join([]string{"CurrentSlot", string(data)}, " ")
}
