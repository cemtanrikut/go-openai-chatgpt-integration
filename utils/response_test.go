package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendJSON(t *testing.T) {
	rec := httptest.NewRecorder()
	data := map[string]string{"message": "Hello, world"}

	SendJSON(rec, data)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", rec.Code)
	}

	var responseBody map[string]string
	err := json.NewDecoder(rec.Body).Decode(&responseBody)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	if responseBody["message"] != "Hello, world" {
		t.Errorf("Expected 'Hello, world', got '%s'", responseBody["message"])
	}
}

func TestSendError(t *testing.T) {
	rec := httptest.NewRecorder()
	errorMessage := "Something went wrong"

	SendError(rec, errorMessage, http.StatusInternalServerError)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code 500, got %d", rec.Code)
	}

	var responseBody map[string]string
	err := json.NewDecoder(rec.Body).Decode(&responseBody)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	if responseBody["error"] != errorMessage {
		t.Errorf("Expected '%s', got '%s'", errorMessage, responseBody["error"])
	}
}
