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

// 存储库修改参数体
type VaultUpdateReq struct {
	Vault *VaultUpdate `json:"vault"`
}

func (o VaultUpdateReq) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"VaultUpdateReq", string(data)}, " ")
}
