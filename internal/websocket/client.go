package websocket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"go-notify-service/internal/models"
	"log"
)

type Client struct {
	UserID int64
	Conn   *websocket.Conn
	Send   chan models.Notification
}

func (c *Client) WritePump() {
	for msg := range c.Send {
		data, _ := json.Marshal(msg)
		if err := c.Conn.WriteMessage(websocket.TextMessage, data); err != nil {
			log.Println("write error:", err)
			return
		}
	}
}
