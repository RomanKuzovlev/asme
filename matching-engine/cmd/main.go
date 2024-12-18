package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/RomanKuzovlev/asme/matching-engine/pkg/api/proto"
	"github.com/RomanKuzovlev/asme/matching-engine/pkg/api/server"
	"github.com/RomanKuzovlev/asme/matching-engine/pkg/matching"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, proceeding with system environment variables")
	}

	tradingPair := os.Getenv("TRADING_PAIR")
	if tradingPair == "" {
		log.Fatalf("TRADING_PAIR environment variable is required")
	}

	engine, err := matching.NewMatchingEngine(tradingPair)
	if err != nil {
		log.Fatalf("Failed to initialize matching engine: %v", err)
	}

	grpcServer := grpc.NewServer()
	matchingServer := server.NewGRPCServer(engine)
	proto.RegisterMatchingEngineServiceServer(grpcServer, matchingServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	log.Println("gRPC server for Matching Engine is running on port 50051...")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down gRPC server...")
	grpcServer.GracefulStop()
	log.Println("Server stopped gracefully.")
}
