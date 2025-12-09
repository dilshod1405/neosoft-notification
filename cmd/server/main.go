package main

import (
	"go-notify-service/internal/http"
	"go-notify-service/internal/redis"
	"go-notify-service/internal/websocket"
	"log"
)

func main() {
	hub := websocket.NewHub()
	go hub.Run()

	// Start Redis Consumer
	consumer := redis.NewStreamConsumer(hub)
	consumer.InitGroup()
	go consumer.Start()

	// HTTP + WS server
	router := http.NewRouter(hub)

	log.Println("🚀 Notification service with Redis Streams running on :8081")
	router.Run(":8081")
}
