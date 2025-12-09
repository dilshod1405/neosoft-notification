package redis

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"strings"

	goredis "github.com/redis/go-redis/v9"
	"go-notify-service/internal/models"
	"go-notify-service/internal/websocket"
)

type StreamConsumer struct {
	Client *goredis.Client
	Hub    *websocket.Hub
}

func NewStreamConsumer(hub *websocket.Hub) *StreamConsumer {
	rdb := goredis.NewClient(&goredis.Options{
		Addr: "localhost:6379",
	})

	return &StreamConsumer{
		Client: rdb,
		Hub:    hub,
	}
}

func (sc *StreamConsumer) InitGroup() {
	ctx := context.Background()
	err := sc.Client.XGroupCreateMkStream(ctx,
		"notifications_stream",
		"notif_group",
		"$",
	).Err()

	if err != nil && !strings.Contains(err.Error(), "BUSYGROUP") {
		log.Println("Error creating consumer group:", err)
	} else {
		log.Println("Redis consumer group OK")
	}
}

func (sc *StreamConsumer) Start() {
	ctx := context.Background()

	for {
		streams, err := sc.Client.XReadGroup(ctx, &goredis.XReadGroupArgs{
			Group:    "notif_group",
			Consumer: "notif_worker_1",
			Streams:  []string{"notifications_stream", ">"},
			Count:    1,
			Block:    0 * time.Second, // Block until message arrives
		}).Result()

		if err != nil {
			log.Println("Redis stream read error:", err)
			continue
		}

		for _, stream := range streams {
			for _, message := range stream.Messages {
				raw := message.Values["data"].(string)

				// Extract message
				var notif models.Notification
				json.Unmarshal([]byte(raw), &notif)

				// Broadcast to WebSocket
				sc.Hub.Broadcast <- notif

				// Acknowledge message
				sc.Client.XAck(ctx, "notifications_stream", "notif_group", message.ID)
			}
		}
	}
}
