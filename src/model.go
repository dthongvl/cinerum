package src

type Message struct {
	Room    string `json:"room"`
	Username string `json:"username"`
	MessageType int `json:"messageType"`
	Message  string `json:"message"`
}
