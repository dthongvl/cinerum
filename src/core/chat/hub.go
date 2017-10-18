package chat

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/dthongvl/cinerum/src/model"
)

var MyHub = NewHub()

type Hub struct {
	rooms map[string]map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan model.Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan model.Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		rooms:      make(map[string]map[*Client]bool),
	}
}

func (h *Hub) Register(client *Client) {
	h.register <- client
}

func (h *Hub) Run() {
	go func() {
		for {
			select {
			case client := <-h.register:
				clients := h.rooms[client.RoomID]
				if clients == nil {
					clients = make(map[*Client]bool)
					h.rooms[client.RoomID] = clients
					log.Info("NEW ROOM: " + client.RoomID)
				}
				h.rooms[client.RoomID][client] = true
			case client := <-h.unregister:
				clients := h.rooms[client.RoomID]
				if clients != nil {
					if _, ok := h.rooms[client.RoomID]; ok {
						delete(h.rooms[client.RoomID], client)
						close(client.Send)
						log.Info("CLIENT LEAVE")
						if len(h.rooms[client.RoomID]) == 0 {
							delete(h.rooms, client.RoomID)
							log.Info("REMOVE ROOM: " + client.RoomID)
						}
					}
				}
			case msg := <-h.broadcast:
				jsonMsg, err := json.Marshal(msg)
				log.Info(msg)
				if err == nil {
					clients := h.rooms[msg.RoomID]
					for client := range clients {
						select {
						case client.Send <- jsonMsg:
						default:
							close(client.Send)
							delete(h.rooms[msg.RoomID], client)
							if len(h.rooms[msg.RoomID]) == 0 {
								delete(h.rooms, msg.RoomID)
							}
						}
					}
				}
			}
		}
	}()
}
