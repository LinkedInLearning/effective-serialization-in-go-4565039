package shop

type OrderRequest struct {
	ClientID string `json:"client_id"`
	SKU      string `json:"sku"`
	Amount   int    `json:"amount"`
}

var missingData = `
{
  "client_id": "007",
  "sku": "laser-pen-2x"
}
`

var zeroData = `
{
  "client_id": "007",
  "sku": "laser-pen-2x",
  "amount": 0
}
`
