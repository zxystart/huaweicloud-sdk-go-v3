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

// Response Object
type CreateVaultTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateVaultTagsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateVaultTagsResponse", string(data)}, " ")
}
