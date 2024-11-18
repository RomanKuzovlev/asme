package server

import (
	"context"

	"github.com/RomanKuzovlev/asme/matching-engine/pkg/api/proto"
	"github.com/RomanKuzovlev/asme/matching-engine/pkg/matching"
	ob "github.com/muzykantov/orderbook"
	"github.com/shopspring/decimal"
)

type GRPCServer struct {
	proto.UnimplementedMatchingEngineServiceServer
	MatchingEngine *matching.MatchingEngine
}

func NewGRPCServer(engine *matching.MatchingEngine) *GRPCServer {
	return &GRPCServer{
		MatchingEngine: engine,
	}
}

func (s *GRPCServer) AddOrder(ctx context.Context, req *proto.AddOrderRequest) (*proto.AddOrderResponse, error) {
	var side ob.Side
	switch req.Side {
	case "buy":
		side = ob.Buy
	case "sell":
		side = ob.Sell
	default:
		return &proto.AddOrderResponse{
			Success: false,
			Message: "Invalid order side",
		}, nil
	}

	quantity := decimal.NewFromFloat(req.Quantity)
	price := decimal.NewFromFloat(req.Price)

	s.MatchingEngine.PlaceOrder(req.OrderId, side, quantity, price)

	return &proto.AddOrderResponse{
		Success: true,
		Message: "Order added successfully",
	}, nil
}
