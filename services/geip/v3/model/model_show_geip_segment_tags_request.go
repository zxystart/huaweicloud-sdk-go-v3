package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGeipSegmentTagsRequest Request Object
type ShowGeipSegmentTagsRequest struct {

	// 全域弹性公网IP的id
	ResourceId string `json:"resource_id"`
}

func (o ShowGeipSegmentTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGeipSegmentTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowGeipSegmentTagsRequest", string(data)}, " ")
}
