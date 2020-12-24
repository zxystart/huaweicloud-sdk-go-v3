/*
 * CBR
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type ResourceExtraInfo struct {
	// 需要排除备份的卷id。仅虚拟机磁盘级备份有效
	ExcludeVolumes *[]string `json:"exclude_volumes,omitempty"`
	// 要备份的卷
	IncludeVolumes *[]ResourceExtraInfoIncludeVolumes `json:"include_volumes,omitempty"`
}

func (o ResourceExtraInfo) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ResourceExtraInfo", string(data)}, " ")
}
