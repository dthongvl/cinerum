package src

import "github.com/gorilla/websocket"

type Message struct {
	Room    string `json:"room"`
	Username string `json:"username"`
	MessageType int `json:"messageType"`
	Message  string `json:"message"`
}

type Room struct {
	Members []websocket.Conn
	TotalMember int64
	RoomId string
}

var rooms []Room
