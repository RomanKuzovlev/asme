package matching

import (
	"fmt"
	"log"

	ob "github.com/muzykantov/orderbook"
	"github.com/shopspring/decimal"
)

// MatchingEngine wraps the external orderbook and provides custom logic
type MatchingEngine struct {
	OrderBook *ob.OrderBook
}

// NewMatchingEngine creates a new MatchingEngine instance
func NewMatchingEngine() *MatchingEngine {
	return &MatchingEngine{
		OrderBook: ob.NewOrderBook(),
	}
}

// PlaceOrder adds an order to the order book
func (me *MatchingEngine) PlaceOrder(orderID string, side ob.Side, quantity, price decimal.Decimal) {
	done, partial, partialQuantityProcessed, err := me.OrderBook.ProcessLimitOrder(side, orderID, quantity, price)
	if err != nil {
		log.Fatalf("Failed to place limit order: %v", err)
	}

	fmt.Println("Done orders:", done)
	if partial != nil {
		fmt.Printf("Partial order processed: %+v, Quantity processed: %s\n", partial, partialQuantityProcessed)
	} else {
		fmt.Println("Order fully processed.")
	}
}

func (me *MatchingEngine) ListOrders() {
	asks, bids := me.OrderBook.Depth()
	fmt.Println("Asks:", asks)
	fmt.Println("Bids:", bids)
}
