package main

import (
	"net/http"
	"testing"
	"time"
)

func TestMainFunction(t *testing.T) {
	// Run main in a goroutine to avoid blocking
	go main()

	// Wait for the server to start
	time.Sleep(1 * time.Second)

	// Make a simple request to the running server
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		t.Fatalf("Failed to send request to the server: %v", err)
	}
	defer resp.Body.Close()

	// Check if the server responded with the correct status
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}
