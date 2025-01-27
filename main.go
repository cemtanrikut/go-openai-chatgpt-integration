package main

import (
	"log"
	"net/http"

	"github.com/cemtanrikut/go-openai-chatgpt-integration/config"
	"github.com/cemtanrikut/go-openai-chatgpt-integration/handler"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Define routes
	http.HandleFunc("/chat", handler.ChatHandler)

	// Start the server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
