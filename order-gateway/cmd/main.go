package main

import (
	"log"
	"net/http"
	"os"

	"github.com/RomanKuzovlev/asme/order-gateway/internal/api"
	"github.com/RomanKuzovlev/asme/order-gateway/internal/gateway"
	"github.com/RomanKuzovlev/asme/order-gateway/internal/integrations"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env file found")
	}

	grpcAddress := os.Getenv("MATCHING_ENGINE_ADDRESS")
	if grpcAddress == "" {
		log.Fatalf("MATCHING_ENGINE_ADDRESS environment variable is required")
	}

	matchingClient, err := integrations.NewMatchingEngineClient(grpcAddress)
	if err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}

	gatewayService := gateway.NewOrderGatewayService(matchingClient)

	orderHandler := api.NewOrderHandler(gatewayService)

	router := mux.NewRouter()
	router.HandleFunc("/orders", orderHandler.CreateOrder).Methods(http.MethodPost)

	port := ":8080"
	log.Printf("Order Gateway is running on port %s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
