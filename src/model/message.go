package model

type MessageBroadcast struct {
	RoomID   string `json:"roomID"`
	Username string `json:"username"`
	Data     string `json:"data"`
}