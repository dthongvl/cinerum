package src

type message struct {
	RoomId    string `json:"roomId"`
	Username string `json:"username"`
	Type string `json:"type"`
	Data  string `json:"data"`
}
