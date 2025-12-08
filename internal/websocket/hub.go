package websocket

import "go-notify-service/internal/models"

type Hub struct {
	Clients    map[int64]map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan models.Notification
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[int64]map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan models.Notification),
	}
}

func (h *Hub) Run() {
	for {
		select {

		case client := <-h.Register:
			if h.Clients[client.UserID] == nil {
				h.Clients[client.UserID] = make(map[*Client]bool)
			}
			h.Clients[client.UserID][client] = true

		case client := <-h.Unregister:
			if clients, ok := h.Clients[client.UserID]; ok {
				delete(clients, client)
				close(client.Send)
			}

		case notif := <-h.Broadcast:
			if clients, ok := h.Clients[notif.UserID]; ok {
				for c := range clients {
					c.Send <- notif
				}
			}
		}
	}
}
