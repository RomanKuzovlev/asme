package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/RomanKuzovlev/asme/order-gateway/internal/gateway"
	"github.com/RomanKuzovlev/asme/order-gateway/internal/models"
)

type OrderHandler struct {
	Service *gateway.OrderGatewayService
}

func NewOrderHandler(service *gateway.OrderGatewayService) *OrderHandler {
	return &OrderHandler{
		Service: service,
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req models.OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	resp, err := h.Service.RouteOrder(r.Context(), req)
	if err != nil {
		log.Printf("Error routing order: %v", err)
		http.Error(w, "Failed to process order", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
