package main

import (
	"log"
	"go-notify-service/internal/http"
	"go-notify-service/internal/websocket"
)

func main() {
	hub := websocket.NewHub()
	go hub.Run()

	router := http.NewRouter(hub)

	log.Println("🚀 Notification service running on :8081")
	err := router.Run(":8081")
	if err != nil {
		log.Fatal(err)
	}
}
