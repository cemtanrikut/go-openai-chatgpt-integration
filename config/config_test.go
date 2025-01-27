package config

import (
	"os"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	// Set up: create a mock .env file
	mockEnv := "OPENAI_API_KEY=test_key"
	err := os.WriteFile(".env", []byte(mockEnv), 0644)
	if err != nil {
		t.Fatalf("Failed to create mock .env file: %v", err)
	}
	defer os.Remove(".env")

	// Test loading environment variables
	LoadEnv()

	if os.Getenv("OPENAI_API_KEY") != "test_key" {
		t.Errorf("Expected OPENAI_API_KEY to be 'test_key', got '%s'", os.Getenv("OPENAI_API_KEY"))
	}
}
