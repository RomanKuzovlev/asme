package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/RomanKuzovlev/asme/order-gateway/internal/api"
	"github.com/RomanKuzovlev/asme/order-gateway/internal/gateway"
	"github.com/RomanKuzovlev/asme/order-gateway/internal/integrations"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	grpcAddress, exists := os.LookupEnv("MATCHING_ENGINE_ADDRESS")
	if !exists || grpcAddress == "" {
		log.Fatal("MATCHING_ENGINE_ADDRESS environment variable is required")
	}

	matchingClient, err := integrations.NewMatchingEngineClient(grpcAddress)
	if err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}

	gatewayService := gateway.NewOrderGatewayService(matchingClient)
	orderHandler := api.NewOrderHandler(gatewayService)

	router := mux.NewRouter()
	router.HandleFunc("/orders", orderHandler.CreateOrder).Methods(http.MethodPost)

	const port = ":8080"
	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	go func() {
		log.Printf("Order Gateway is running on port %s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped gracefully")
}
