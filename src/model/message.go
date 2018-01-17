package model

type MessageBroadcast struct {
	RoomId   string `json:"roomID"`
	Username string `json:"username"`
	Data     string `json:"data"`
	Type     string `json:"type"`
}
