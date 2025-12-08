package http

import (
	"github.com/gin-gonic/gin"
	"go-notify-service/internal/websocket"
)

func NewRouter(hub *websocket.Hub) *gin.Engine {
	r := gin.Default()

	r.GET("/ws", func(c *gin.Context) {
		websocket.HandleWebSocket(hub, c)
	})

	r.POST("/publish", func(c *gin.Context) {
		PublishNotification(hub, c)
	})

	return r
}
