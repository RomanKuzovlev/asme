package test

import (
	"context"
	"testing"

	"log"
	"net"

	"github.com/RomanKuzovlev/asme/matching-engine/pkg/api/proto"
	"github.com/RomanKuzovlev/asme/matching-engine/pkg/api/server"
	"github.com/RomanKuzovlev/asme/matching-engine/pkg/matching"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	grpcServer := grpc.NewServer()
	matchingEngine := matching.NewMatchingEngine()
	proto.RegisterMatchingEngineServiceServer(grpcServer, server.NewGRPCServer(matchingEngine))
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestAddOrder(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	require.NoError(t, err)
	defer conn.Close()

	client := proto.NewMatchingEngineServiceClient(conn)
	resp, err := client.AddOrder(ctx, &proto.AddOrderRequest{
		OrderId:  "test-order-1",
		Side:     "buy",
		Quantity: 10.0,
		Price:    100.0,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.True(t, resp.Success)
	require.Equal(t, "Order added successfully", resp.Message)
}
