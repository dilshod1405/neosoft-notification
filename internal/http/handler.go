package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-notify-service/internal/models"
	"go-notify-service/internal/websocket"
)

func PublishNotification(hub *websocket.Hub, c *gin.Context) {
	var notif models.Notification

	if err := c.ShouldBindJSON(&notif); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	// Hub ga yuborish (broadcast)
	hub.Broadcast <- notif

	c.JSON(http.StatusOK, gin.H{"status": "sent"})
}
