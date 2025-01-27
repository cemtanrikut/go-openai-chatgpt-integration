package handler

import (
	"encoding/json"
	"net/http"
)

type ChatHandler struct {
	GetChatResponseFunc func(prompt string) (string, error)
}

func (h ChatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var chatReq struct {
		Prompt string `json:"prompt"`
	}
	err := json.NewDecoder(r.Body).Decode(&chatReq)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	response, err := h.GetChatResponseFunc(chatReq.Prompt)
	if err != nil {
		http.Error(w, "Error communicating with OpenAI", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"response": response})
}
