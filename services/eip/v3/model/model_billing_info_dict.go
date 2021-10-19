package model

import (
	"encoding/json"

	"strings"
)

// 订单信息, 有订单表示包周期
type BillingInfoDict struct {
	// 订单信息

	OrderId *string `json:"order_id,omitempty"`
	// 产品id

	ProductId *string `json:"product_id,omitempty"`
}

func (o BillingInfoDict) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BillingInfoDict struct{}"
	}

	return strings.Join([]string{"BillingInfoDict", string(data)}, " ")
}
