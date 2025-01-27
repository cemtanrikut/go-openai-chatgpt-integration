package model

type ChatRequest struct {
	Prompt string `json:"prompt" validate:"required"`
}

// Add a method to the ChatRequest struct
func (c *ChatRequest) IsValid() bool {
	return c.Prompt != ""
}
