package integrations

import (
	"context"
	"log"
	"time"

	"github.com/RomanKuzovlev/asme/matching-engine/pkg/api/proto"
	"google.golang.org/grpc"
)

type MatchingEngineClient struct {
	client proto.MatchingEngineServiceClient
}

func NewMatchingEngineClient(address string) (*MatchingEngineClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		return nil, err
	}

	return &MatchingEngineClient{
		client: proto.NewMatchingEngineServiceClient(conn),
	}, nil
}

func (m *MatchingEngineClient) AddOrder(ctx context.Context, orderID, side string, quantity, price float64, pair string) (*proto.AddOrderResponse, error) {
	req := &proto.AddOrderRequest{
		OrderId:  orderID,
		Side:     side,
		Quantity: quantity,
		Price:    price,
		Pair:     pair,
	}

	resp, err := m.client.AddOrder(ctx, req)
	if err != nil {
		log.Printf("Error calling AddOrder: %v", err)
		return nil, err
	}

	return resp, nil
}
