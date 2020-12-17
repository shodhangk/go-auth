package jsons

type CartItemJson struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Quantity     uint   `json:"quantity"`
	PricePerUnit int    `json:"price_per_unit"`
}

type CartJson struct {
	Items []CartItemJson `json:"items"`
}
