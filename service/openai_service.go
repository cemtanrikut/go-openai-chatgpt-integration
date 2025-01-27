package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

type OpenAIRequest struct {
	Model    string              `json:"model"`
	Messages []map[string]string `json:"messages"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func GetChatResponse(prompt string) (string, error) {
	// Fetch API URL and API Key from environment variables
	apiURL := os.Getenv("OPENAI_API_URL")
	if apiURL == "" {
		return "", errors.New("OPENAI_API_URL is not set")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", errors.New("OPENAI_API_KEY is not set")
	}

	// Prepare the OpenAI API request payload
	requestBody, err := json.Marshal(OpenAIRequest{
		Model: "gpt-4",
		Messages: []map[string]string{
			{"role": "user", "content": prompt},
		},
	})
	if err != nil {
		return "", err
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Perform the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("OpenAI API error")
	}

	// Parse the OpenAI API response
	var openAIResp OpenAIResponse
	err = json.NewDecoder(resp.Body).Decode(&openAIResp)
	if err != nil {
		return "", err
	}

	// Return the response content
	return openAIResp.Choices[0].Message.Content, nil
}
