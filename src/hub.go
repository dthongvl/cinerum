package src

import (
	"encoding/json"
	"log"
)

type Hub struct {
	rooms map[string]map[*Client]bool

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
		rooms:    make(map[string]map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			clients := h.rooms[client.roomId]
			if clients == nil {
				clients = make(map[*Client]bool)
				h.rooms[client.roomId] = clients
				log.Println("NEW ROOM: " + client.roomId)
			}
			h.rooms[client.roomId][client] = true
		case client := <-h.unregister:
			clients := h.rooms[client.roomId]
			if clients != nil {
				if _, ok := h.rooms[client.roomId]; ok {
					delete(h.rooms[client.roomId], client)
					close(client.send)
					log.Println("CLIENT LEAVE")
					if len(h.rooms[client.roomId]) == 0 {
						delete(h.rooms, client.roomId)
						log.Println("DELETE ROOM: " + client.roomId)
					}
				}
			}
		case msg := <-h.broadcast:
			jsonMsg, err := json.Marshal(msg)
			log.Println(msg)
			if err == nil {
				clients := h.rooms[msg.RoomId]
				for client := range clients {
					select {
					case client.send <- jsonMsg:
					default:
						close(client.send)
						delete(h.rooms[msg.RoomId], client)
						if len(h.rooms[msg.RoomId]) == 0 {
							delete(h.rooms, msg.RoomId)
						}
					}
				}
			}
		}
	}
}
