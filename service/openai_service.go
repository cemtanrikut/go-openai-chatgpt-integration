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
	openAIReq := OpenAIRequest{
		Model: "gpt-4",
		Messages: []map[string]string{
			{"role": "user", "content": prompt},
		},
	}

	reqBody, err := json.Marshal(openAIReq)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", os.Getenv("OPENAI_API_URL"), bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("OpenAI API error")
	}

	var openAIResp OpenAIResponse
	err = json.NewDecoder(resp.Body).Decode(&openAIResp)
	if err != nil {
		return "", err
	}

	return openAIResp.Choices[0].Message.Content, nil
}
