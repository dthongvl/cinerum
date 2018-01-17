package chat

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/dthongvl/cinerum/src/model"
	"strconv"
)

type Hub struct {
	rooms map[string]map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan model.MessageBroadcast

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan model.MessageBroadcast),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		rooms:      make(map[string]map[*Client]bool),
	}
}

func (h *Hub) Register(client *Client) {
	h.register <- client
}

func (h *Hub) GetTotalOnline(roomId string) int {
	if _, ok := h.rooms[roomId]; ok {
		return len(h.rooms[roomId])
	}
	return 0
}

func (h *Hub) Run() {
	go func() {
		for {
			select {
			case client := <-h.register:
				clients := h.rooms[client.RoomId]
				if clients == nil {
					clients = make(map[*Client]bool)
					h.rooms[client.RoomId] = clients
					log.Info("NEW ROOM: " + client.RoomId)
				}
				h.rooms[client.RoomId][client] = true
				go h.broadcastTotalOnline(client.RoomId, client.Username)
			case client := <-h.unregister:
				if _, ok := h.rooms[client.RoomId]; ok {
					delete(h.rooms[client.RoomId], client)
					close(client.Send)
					log.Info("CLIENT LEAVE")
					if len(h.rooms[client.RoomId]) == 0 {
						delete(h.rooms, client.RoomId)
						log.Info("REMOVE ROOM: " + client.RoomId)
					}
					go h.broadcastTotalOnline(client.RoomId, client.Username)
				}
			case msg := <-h.broadcast:
				jsonMsg, err := json.Marshal(msg)
				log.Println("BROADCAST:", msg, "TO ROOM", msg.RoomId)
				if err == nil {
					clients := h.rooms[msg.RoomId]
					for client := range clients {
						select {
						case client.Send <- jsonMsg:
						default:
							close(client.Send)
							delete(h.rooms[msg.RoomId], client)
							if len(h.rooms[msg.RoomId]) == 0 {
								delete(h.rooms, msg.RoomId)
							}
						}
					}
				} else {
					log.Println("ERROR")
				}
			}
		}
	}()
}

func (h *Hub) broadcastTotalOnline(roomId string, username string) {
	totalOnline := h.GetTotalOnline(roomId)
	msg := model.MessageBroadcast{
		RoomId:   roomId,
		Username: username,
		Data:     strconv.Itoa(totalOnline),
		Type:     "online",
	}
	h.broadcast <- msg
}
