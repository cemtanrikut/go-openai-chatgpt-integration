package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestChatHandler(t *testing.T) {
	// Mock GetChatResponse function for successful case
	mockGetChatResponse := func(prompt string) (string, error) {
		return "Mock response", nil
	}

	// Mock GetChatResponse function for error case
	mockGetChatResponseError := func(prompt string) (string, error) {
		return "", errors.New("mock error")
	}

	tests := []struct {
		name               string
		mockFunc           func(string) (string, error)
		requestBody        string
		expectedStatusCode int
		expectedBody       interface{}
	}{
		{
			name:               "Valid request",
			mockFunc:           mockGetChatResponse,
			requestBody:        `{"prompt": "Test prompt"}`,
			expectedStatusCode: http.StatusOK,
			expectedBody:       map[string]interface{}{"response": "Mock response"},
		},
		{
			name:               "Invalid JSON in request body",
			mockFunc:           mockGetChatResponse,
			requestBody:        `invalid-json`,
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "Invalid request\n",
		},
		{
			name:               "GetChatResponse returns an error",
			mockFunc:           mockGetChatResponseError,
			requestBody:        `{"prompt": "Test prompt"}`,
			expectedStatusCode: http.StatusInternalServerError,
			expectedBody:       "Error communicating with OpenAI\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new handler instance with the mock function
			handler := ChatHandler{
				GetChatResponseFunc: tc.mockFunc,
			}

			// Create a request
			req := httptest.NewRequest(http.MethodPost, "/chat", bytes.NewBuffer([]byte(tc.requestBody)))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			// Call the handler
			handler.ServeHTTP(rec, req)

			// Check the status code
			if rec.Code != tc.expectedStatusCode {
				t.Errorf("Expected status code %d, got %d", tc.expectedStatusCode, rec.Code)
			}

			// Check the response body
			if expectedMap, ok := tc.expectedBody.(map[string]interface{}); ok {
				// For JSON responses
				var actualBody map[string]interface{}
				err := json.NewDecoder(rec.Body).Decode(&actualBody)
				if err != nil {
					t.Fatalf("Failed to decode response body: %v", err)
				}
				if !reflect.DeepEqual(actualBody, expectedMap) {
					t.Errorf("Expected body %v, got %v", expectedMap, actualBody)
				}
			} else if expectedString, ok := tc.expectedBody.(string); ok {
				// For plain text responses
				actualBody := rec.Body.String()
				if actualBody != expectedString {
					t.Errorf("Expected body %q, got %q", expectedString, actualBody)
				}
			}
		})
	}
}
