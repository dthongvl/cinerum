package model

type Message struct {
	RoomID   string `json:"roomID"`
	Username string `json:"username"`
	Type     string `json:"type"`
	Data     string `json:"data"`
}
