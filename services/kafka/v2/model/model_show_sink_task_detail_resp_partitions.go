package model

import (
	"encoding/json"

	"strings"
)

type ShowSinkTaskDetailRespPartitions struct {
	// 分区ID。

	PartitionId *string `json:"partition_id,omitempty"`
	// 运行状态。

	Status *string `json:"status,omitempty"`
	// 已转储的消息偏移量。

	LastTransferOffset *string `json:"last_transfer_offset,omitempty"`
	// 消息偏移量。

	LogEndOffset *string `json:"log_end_offset,omitempty"`
	// 积压的消息数。

	Lag *string `json:"lag,omitempty"`
}

func (o ShowSinkTaskDetailRespPartitions) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowSinkTaskDetailRespPartitions struct{}"
	}

	return strings.Join([]string{"ShowSinkTaskDetailRespPartitions", string(data)}, " ")
}
