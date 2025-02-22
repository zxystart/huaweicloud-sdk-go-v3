package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidatePolicyResponse Response Object
type ValidatePolicyResponse struct {
	Findings *[]ValidatePolicyFinding `json:"findings,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ValidatePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidatePolicyResponse struct{}"
	}

	return strings.Join([]string{"ValidatePolicyResponse", string(data)}, " ")
}
