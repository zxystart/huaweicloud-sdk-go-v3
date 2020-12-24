/*
 * AS
 *
 * 弹性伸缩API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowScalingGroupRequest struct {
	ScalingGroupId string `json:"scaling_group_id"`
}

func (o ShowScalingGroupRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowScalingGroupRequest", string(data)}, " ")
}
