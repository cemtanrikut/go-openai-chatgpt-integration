package model

import (
	"encoding/json"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestChatRequestSerialization(t *testing.T) {
	// Test JSON serialization
	input := ChatRequest{Prompt: "Test prompt"}
	expectedJSON := `{"prompt":"Test prompt"}`

	data, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("Failed to serialize ChatRequest: %v", err)
	}

	actualJSON := string(data)
	if actualJSON != expectedJSON {
		t.Errorf("Expected JSON: %s, got: %s", expectedJSON, actualJSON)
	}
}

func TestChatRequestDeserialization(t *testing.T) {
	// Test JSON deserialization
	inputJSON := `{"prompt":"Test prompt"}`
	expected := ChatRequest{Prompt: "Test prompt"}

	var actual ChatRequest
	err := json.Unmarshal([]byte(inputJSON), &actual)
	if err != nil {
		t.Fatalf("Failed to deserialize ChatRequest: %v", err)
	}

	if actual.Prompt != expected.Prompt {
		t.Errorf("Expected prompt: %s, got: %s", expected.Prompt, actual.Prompt)
	}
}

func TestChatRequestValidation(t *testing.T) {
	// Validator instance
	validate := validator.New()

	// Valid case
	validRequest := ChatRequest{Prompt: "Valid prompt"}
	err := validate.Struct(validRequest)
	if err != nil {
		t.Errorf("Expected no validation errors, but got: %v", err)
	}

	// Invalid case: Missing Prompt
	invalidRequest := ChatRequest{} // Prompt is empty
	err = validate.Struct(invalidRequest)
	if err == nil {
		t.Error("Expected validation error, but got none")
	} else {
		// Check the specific validation error
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			t.Fatalf("Expected validator.ValidationErrors, got: %v", err)
		}

		// Verify that the error is for the "Prompt" field
		for _, fieldError := range validationErrors {
			if fieldError.Field() != "Prompt" || fieldError.Tag() != "required" {
				t.Errorf("Unexpected validation error: Field=%s, Tag=%s", fieldError.Field(), fieldError.Tag())
			}
		}
	}
}
