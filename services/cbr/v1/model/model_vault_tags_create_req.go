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

//
type VaultTagsCreateReq struct {
	Tag *Tag `json:"tag,omitempty"`
}

func (o VaultTagsCreateReq) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"VaultTagsCreateReq", string(data)}, " ")
}
