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

// Request Object
type ShowVaultProjectTagRequest struct {
}

func (o ShowVaultProjectTagRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowVaultProjectTagRequest", string(data)}, " ")
}
