package matching

import (
	"fmt"
	"log"

	ob "github.com/muzykantov/orderbook"
	"github.com/shopspring/decimal"
)

// MatchingEngine wraps the external orderbook and provides custom logic for a specific trading pair.
type MatchingEngine struct {
	OrderBook   *ob.OrderBook
	TradingPair string
}

// NewMatchingEngine creates a new MatchingEngine instance for the specified trading pair.
func NewMatchingEngine(tradingPair string) (*MatchingEngine, error) {
	if tradingPair == "" {
		return nil, fmt.Errorf("trading pair is not set for the engine")
	}
	return &MatchingEngine{
		OrderBook:   ob.NewOrderBook(),
		TradingPair: tradingPair,
	}, nil
}

// PlaceOrder adds an order to the order book.
func (me *MatchingEngine) PlaceOrder(orderID string, side ob.Side, quantity, price decimal.Decimal) error {
	done, partial, partialQuantityProcessed, err := me.OrderBook.ProcessLimitOrder(side, orderID, quantity, price)
	if err != nil {
		log.Printf("Error placing limit order for trading pair %s: %v", me.TradingPair, err)
		return fmt.Errorf("failed to place limit order: %w", err)
	}

	log.Printf("Order placed in trading pair %s: OrderID=%s, Side=%v, Quantity=%s, Price=%s",
		me.TradingPair, orderID, side, quantity, price)

	log.Printf("Done orders for trading pair %s: %v", me.TradingPair, done)
	if partial != nil {
		log.Printf("Partial order processed for trading pair %s: %+v, Quantity processed: %s",
			me.TradingPair, partial, partialQuantityProcessed)
	} else {
		log.Printf("Order fully processed for trading pair %s.", me.TradingPair)
	}

	return nil
}

// ListOrders lists the current asks and bids in the order book.
func (me *MatchingEngine) ListOrders() {
	asks, bids := me.OrderBook.Depth()

	log.Printf("Order book depth for trading pair %s:", me.TradingPair)
	log.Printf("Asks: %s", formatPriceLevels(asks))
	log.Printf("Bids: %s", formatPriceLevels(bids))
}

// formatPriceLevels formats price levels into a readable string representation.
func formatPriceLevels(levels []*ob.PriceLevel) string {
	if len(levels) == 0 {
		return "None"
	}
	var result string
	for _, level := range levels {
		result += fmt.Sprintf("[Price: %s, Quantity: %s] ", level.Price, level.Quantity)
	}
	return result
}
