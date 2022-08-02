package main

import (
	"fmt"
	"log"
	"net/http"

	"server/handlers"
)

func main() {
	http.HandleFunc("/ws", handlers.NewWebSocketHandler().Handle)

	port := "8080"
	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Panicln("Error starting server:", err)
	}
}
