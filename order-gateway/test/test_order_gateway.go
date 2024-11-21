package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	url := "http://localhost:8080/orders"

	orderPayload := `{
		"order_id": "12345",
		"side": "buy",
		"quantity": 10,
		"price": 5000.00,
		"pair": "BTC/USD"
	}`

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(orderPayload)))
	if err != nil {
		log.Fatalf("Failed to create HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Printf("Response Status: %s\n", resp.Status)
	fmt.Println("Response Body:")
	if err := resp.Write(&bytes.Buffer{}); err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}
}
