package main

import (
	"log"
	"net/http"

	"github.com/cemtanrikut/go-openai-chatgpt-integration/config"
	"github.com/cemtanrikut/go-openai-chatgpt-integration/handler"
	"github.com/cemtanrikut/go-openai-chatgpt-integration/service"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Create an instance of ChatHandler with the actual GetChatResponse function
	chatHandler := handler.ChatHandler{
		GetChatResponseFunc: service.GetChatResponse,
	}

	// Define routes
	http.Handle("/chat", chatHandler)

	// Start the server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
