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

// 镜像元素
type ImageData struct {
	//
	ImageId *string `json:"image_id,omitempty"`
}

func (o ImageData) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ImageData", string(data)}, " ")
}
