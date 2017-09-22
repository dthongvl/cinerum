package src

import "encoding/json"

type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case msg := <-h.broadcast:
			jsonMsg, err := json.Marshal(msg)
			if err == nil {
				for client := range h.clients {
					select {
					case client.send <- jsonMsg:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}
		}
	}
}