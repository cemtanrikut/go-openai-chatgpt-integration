package service

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetChatResponse(t *testing.T) {
	t.Run("Missing OPENAI_API_URL", func(t *testing.T) {
		os.Unsetenv("OPENAI_API_URL")
		os.Setenv("OPENAI_API_KEY", "mock_key")
		defer os.Unsetenv("OPENAI_API_KEY")

		_, err := GetChatResponse("Test prompt")
		if err == nil || err.Error() != "OPENAI_API_URL is not set" {
			t.Fatalf("Expected error for missing OPENAI_API_URL, got: %v", err)
		}
	})

	t.Run("Missing OPENAI_API_KEY", func(t *testing.T) {
		os.Setenv("OPENAI_API_URL", "http://example.com")
		os.Unsetenv("OPENAI_API_KEY")

		_, err := GetChatResponse("Test prompt")
		if err == nil || err.Error() != "OPENAI_API_KEY is not set" {
			t.Fatalf("Expected error for missing OPENAI_API_KEY, got: %v", err)
		}
	})

	t.Run("Invalid API URL", func(t *testing.T) {
		os.Setenv("OPENAI_API_URL", "://invalid-url")
		os.Setenv("OPENAI_API_KEY", "mock_key")
		defer os.Unsetenv("OPENAI_API_URL")
		defer os.Unsetenv("OPENAI_API_KEY")

		_, err := GetChatResponse("Test prompt")
		if err == nil || err.Error() == "" {
			t.Fatalf("Expected error for invalid API URL, got: %v", err)
		}
	})

	t.Run("Invalid API Response", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{invalid-json}`))
		}))
		defer mockServer.Close()

		os.Setenv("OPENAI_API_URL", mockServer.URL)
		os.Setenv("OPENAI_API_KEY", "mock_key")
		defer os.Unsetenv("OPENAI_API_URL")
		defer os.Unsetenv("OPENAI_API_KEY")

		_, err := GetChatResponse("Test prompt")
		if err == nil || err.Error() == "" {
			t.Fatalf("Expected error for invalid API response, got: %v", err)
		}
	})

	t.Run("Valid API Response", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"choices":[{"message":{"content":"Mock response"}}]}`))
		}))
		defer mockServer.Close()

		os.Setenv("OPENAI_API_URL", mockServer.URL)
		os.Setenv("OPENAI_API_KEY", "mock_key")
		defer os.Unsetenv("OPENAI_API_URL")
		defer os.Unsetenv("OPENAI_API_KEY")

		response, err := GetChatResponse("Test prompt")
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		expectedResponse := "Mock response"
		if response != expectedResponse {
			t.Errorf("Expected response: %q, got: %q", expectedResponse, response)
		}
	})
}
