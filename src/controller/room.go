package controller

import (
	"bytes"
	"encoding/base64"
	"net/http"

	"github.com/dthongvl/cinerum/src/core/chat"
	"github.com/dthongvl/cinerum/template"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func JoinRoom(c echo.Context) error {
	buffer := new(bytes.Buffer)
	template.JoinRoom(c.Param("roomID"), buffer)
	return c.HTML(http.StatusOK, buffer.String())
}

func CreateRoom(c echo.Context) error {
	clientIP := c.RealIP()
	roomID := base64.StdEncoding.EncodeToString([]byte(clientIP + "TMT"))
	return c.Redirect(http.StatusMovedPermanently, "/room/"+roomID)
}

func ServeWebSocket(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("upgrade:", err)
		return err
	}

	client := &chat.Client{ChatHub: chat.MyHub, Conn: conn, Send: make(chan []byte, 256), RoomID: c.QueryParam("roomID")}
	client.ChatHub.Register(client)

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
	return nil
}
