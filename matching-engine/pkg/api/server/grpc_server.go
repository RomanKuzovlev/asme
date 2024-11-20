package server

import (
	"context"
	"fmt"
	"log"

	"github.com/RomanKuzovlev/asme/matching-engine/pkg/api/proto"
	"github.com/RomanKuzovlev/asme/matching-engine/pkg/matching"
	ob "github.com/muzykantov/orderbook"
	"github.com/shopspring/decimal"
)

type GRPCServer struct {
	proto.UnimplementedMatchingEngineServiceServer
	MatchingEngine *matching.MatchingEngine
}

// NewGRPCServer initializes a new gRPC server with a matching engine instance.
func NewGRPCServer(engine *matching.MatchingEngine) *GRPCServer {
	return &GRPCServer{
		MatchingEngine: engine,
	}
}

// AddOrder handles the gRPC request to add an order to the matching engine.
func (s *GRPCServer) AddOrder(ctx context.Context, req *proto.AddOrderRequest) (*proto.AddOrderResponse, error) {
	log.Printf("Received AddOrder request: OrderID=%s, Side=%s, Quantity=%f, Price=%f, Pair=%s",
		req.OrderId, req.Side, req.Quantity, req.Price, req.Pair)

	if req.Pair != s.MatchingEngine.TradingPair {
		err := fmt.Errorf("invalid trading pair: expected %s, got %s", s.MatchingEngine.TradingPair, req.Pair)
		log.Printf("Failed to add order: %v", err)
		return &proto.AddOrderResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	var side ob.Side
	switch req.Side {
	case "buy":
		side = ob.Buy
	case "sell":
		side = ob.Sell
	default:
		err := fmt.Errorf("invalid order side: %s", req.Side)
		log.Printf("Failed to add order: %v", err)
		return &proto.AddOrderResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	quantity := decimal.NewFromFloat(req.Quantity)
	price := decimal.NewFromFloat(req.Price)
	if quantity.Sign() <= 0 {
		err := fmt.Errorf("invalid quantity: must be greater than zero")
		log.Printf("Failed to add order: %v", err)
		return &proto.AddOrderResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}
	if price.Sign() <= 0 {
		err := fmt.Errorf("invalid price: must be greater than zero")
		log.Printf("Failed to add order: %v", err)
		return &proto.AddOrderResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	err := s.MatchingEngine.PlaceOrder(req.OrderId, side, quantity, price)
	if err != nil {
		log.Printf("Failed to add order: %v", err)
		return &proto.AddOrderResponse{
			Success: false,
			Message: fmt.Sprintf("Failed to add order: %v", err),
		}, nil
	}

	log.Printf("Successfully added order: OrderID=%s, Side=%s, Quantity=%f, Price=%f, Pair=%s",
		req.OrderId, req.Side, req.Quantity, req.Price, req.Pair)
	return &proto.AddOrderResponse{
		Success: true,
		Message: "Order added successfully",
	}, nil
}
