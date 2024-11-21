package gateway

import (
	"context"
	"errors"

	api "github.com/RomanKuzovlev/asme/order-gateway/internal/integrations"
	"github.com/RomanKuzovlev/asme/order-gateway/internal/models"
)

type OrderGatewayService struct {
	MatchingClient *api.MatchingEngineClient
}

func NewOrderGatewayService(client *api.MatchingEngineClient) *OrderGatewayService {
	return &OrderGatewayService{
		MatchingClient: client,
	}
}

func (s *OrderGatewayService) RouteOrder(ctx context.Context, req models.OrderRequest) (models.OrderResponse, error) {
	if req.OrderID == "" || req.Side == "" || req.Quantity <= 0 || req.Price <= 0 || req.Pair == "" {
		return models.OrderResponse{}, errors.New("invalid order request")
	}

	resp, err := s.MatchingClient.AddOrder(ctx, req.OrderID, req.Side, req.Quantity, req.Price, req.Pair)
	if err != nil {
		return models.OrderResponse{
			Success: false,
			Message: "Failed to route order to matching engine",
		}, err
	}

	return models.OrderResponse{
		Success: resp.Success,
		Message: resp.Message,
	}, nil
}
