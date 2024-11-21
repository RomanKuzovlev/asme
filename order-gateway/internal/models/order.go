package models

type OrderRequest struct {
	OrderID  string  `json:"order_id"`
	Side     string  `json:"side"` // "buy" or "sell"
	Quantity float64 `json:"quantity"`
	Price    float64 `json:"price"`
	Pair     string  `json:"pair"` // trading pair, e.g., "BTC/USD"
}

type OrderResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
