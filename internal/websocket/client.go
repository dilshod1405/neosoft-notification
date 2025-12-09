package websocket

import (
	"log"
	"time"
	"go-notify-service/internal/models"

	

	"github.com/gorilla/websocket"
)

type Client struct {
    UserID int64
    Conn   *websocket.Conn
    Send   chan models.Notification
}

func (c *Client) WritePump() {
    ticker := time.NewTicker(30 * time.Second)
    defer func() {
        ticker.Stop()
        c.Conn.Close()
    }()

    for {
        select {
        case message, ok := <-c.Send:
            if !ok {
                c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
                return
            }

            c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
            if err := c.Conn.WriteJSON(message); err != nil {
                log.Println("WritePump error:", err)
                return
            }

        case <-ticker.C:
            c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
            if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
                log.Println("Ping error:", err)
                return
            }
        }
    }
}
