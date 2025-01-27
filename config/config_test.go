package config

import (
	"bytes"
	"log"
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

func TestLoadEnv_NoEnvFile(t *testing.T) {
	// Redirect log output to a buffer
	var logBuffer bytes.Buffer
	log.SetOutput(&logBuffer)
	log.SetFlags(0) // Disable timestamp in log
	defer func() {
		log.SetOutput(os.Stderr)    // Restore default log output
		log.SetFlags(log.LstdFlags) // Restore default log flags
	}()

	// Ensure no .env file exists
	os.Remove(".env")

	// Call LoadEnv
	LoadEnv()

	// Check if the log message is written
	logOutput := logBuffer.String()
	expectedLog := "No .env file found. Using system environment variables.\n"
	if logOutput != expectedLog {
		t.Errorf("Expected log output: %q, got: %q", expectedLog, logOutput)
	}
}
