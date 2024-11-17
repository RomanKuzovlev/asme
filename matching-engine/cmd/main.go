package main

import (
	"github.com/RomanKuzovlev/asme/matching-engine/pkg/matching"
	ob "github.com/muzykantov/orderbook"
	"github.com/shopspring/decimal"
)

func main() {
	engine := matching.NewMatchingEngine()

	orderID := "order-1"
	side := ob.Buy
	quantity := decimal.NewFromFloat(1.5)
	price := decimal.NewFromFloat(1000.0)

	engine.PlaceOrder(orderID, side, quantity, price)

	engine.ListOrders()
}
