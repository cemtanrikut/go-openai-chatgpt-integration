package handler

import (
	"encoding/json"
	"net/http"

	"github.com/cemtanrikut/go-openai-chatgpt-integration/model"
	"github.com/cemtanrikut/go-openai-chatgpt-integration/service"
	"github.com/cemtanrikut/go-openai-chatgpt-integration/utils"
)

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	var chatReq model.ChatRequest
	err := json.NewDecoder(r.Body).Decode(&chatReq)
	if err != nil {
		utils.SendError(w, "Invalid request", http.StatusBadRequest)
		return
	}

	response, err := service.GetChatResponse(chatReq.Prompt)
	if err != nil {
		utils.SendError(w, "Error communicating with OpenAI", http.StatusInternalServerError)
		return
	}

	utils.SendJSON(w, map[string]string{"response": response})
}
