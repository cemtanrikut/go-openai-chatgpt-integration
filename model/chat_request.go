package model

type ChatRequest struct {
	Prompt string `json:"prompt" validate:"required"`
}
