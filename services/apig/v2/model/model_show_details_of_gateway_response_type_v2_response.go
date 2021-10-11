package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowDetailsOfGatewayResponseTypeV2Response struct {
	Body           map[string]ResponseInfoResp `json:"body,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowDetailsOfGatewayResponseTypeV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDetailsOfGatewayResponseTypeV2Response struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfGatewayResponseTypeV2Response", string(data)}, " ")
}
