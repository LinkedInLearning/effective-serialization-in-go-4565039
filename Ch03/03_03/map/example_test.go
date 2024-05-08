package shop

import "fmt"

func ExampleParseOrderRequest() {
	data := []byte(`{"client_id": "007", "sku": "laser-pen-2x"}`)
	req, err := ParseOrderRequest(data)
	if err != nil {
		fmt.Println("error (missing):", err)
		return
	}
	fmt.Printf("%s %s %d\n", req.ClientID, req.SKU, req.Amount)

	data = []byte(`{"client_id": "007", "sku": "laser-pen-2x", "amount": 0}`)
	_, err = ParseOrderRequest(data)
	fmt.Println("error:", err)

	// Output:
	// 007 laser-pen-2x 1
	// error: invalid amount: 0
}
