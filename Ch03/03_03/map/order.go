package shop

import (
	"encoding/json"
	"fmt"
)

type OrderRequest struct {
	ClientID string `json:"client_id"`
	SKU      string `json:"sku"`
	Amount   int    `json:"amount"`
}

func ParseOrderRequest(data []byte) (OrderRequest, error) {
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		return OrderRequest{}, err
	}

	req := OrderRequest{
		ClientID: m["client_id"].(string),
		SKU:      m["sku"].(string),
	}

	if val, ok := m["amount"]; ok {
		amount := int(val.(float64))
		if amount <= 0 {
			return OrderRequest{}, fmt.Errorf("invalid amount: %d", amount)
		}
		req.Amount = amount
	} else {
		req.Amount = 1
	}

	return req, nil
}
