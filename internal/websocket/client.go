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
    Closed bool
}


func (c *Client) WritePump(hub *Hub) {
    ticker := time.NewTicker(30 * time.Second)
    defer func() {
        ticker.Stop()
        c.Conn.Close()
    }()

    for {
        select {
        case msg, ok := <-c.Send:
            if !ok {
                return
            }

            c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
            if err := c.Conn.WriteJSON(msg); err != nil {
                hub.Unregister <- c
                return
            }

        case <-ticker.C:
            c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
            if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
                hub.Unregister <- c
                return
            }
        }
    }
}





func (c *Client) ReadPump(hub *Hub) {
    defer func() {
        hub.Unregister <- c
        c.Conn.Close()
    }()

    c.Conn.SetReadLimit(512)
    c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
    c.Conn.SetPongHandler(func(string) error {
        c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
        return nil
    })

    for {
        if _, _, err := c.Conn.ReadMessage(); err != nil {
            log.Println("ReadPump error:", err)
            return
        }
    }
}
