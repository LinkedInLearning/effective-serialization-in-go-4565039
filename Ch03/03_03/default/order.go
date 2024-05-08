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
	req := OrderRequest{
		Amount: 1,
	}
	if err := json.Unmarshal(data, &req); err != nil {
		return OrderRequest{}, err
	}

	if req.Amount <= 0 {
		return OrderRequest{}, fmt.Errorf("invalid amount: %d", req.Amount)
	}

	return req, nil
}
