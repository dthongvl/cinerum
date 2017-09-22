package src

type message struct {
	RoomId    string `json:"roomid"`
	Username string `json:"username"`
	Type string `json:"type"`
	Data  string `json:"data"`
}

type ABC struct {
	RoomId    string `json:"roomid"`
	Username string `json:"username"`
	Type string `json:"type"`
	Data  string `json:"data"`
}