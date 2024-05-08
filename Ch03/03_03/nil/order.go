package shop

import (
	"encoding/json"
	"fmt"
)

type OrderRequest struct {
	ClientID string `json:"client_id"`
	SKU      string `json:"sku"`
	Amount   *int   `json:"amount"`
}

func ParseOrderRequest(data []byte) (OrderRequest, error) {
	var req OrderRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return OrderRequest{}, err
	}

	if req.Amount != nil && *req.Amount <= 0 {
		return OrderRequest{}, fmt.Errorf("invalid amount: %d", *req.Amount)
	}

	if req.Amount == nil {
		amount := 1
		req.Amount = &amount
	}

	return req, nil
}
