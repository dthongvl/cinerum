package src

import (
	"github.com/labstack/echo"
	"net/http"
	"encoding/base64"
	"log"
	"cinerum/template"
	"bytes"
)

func index(c echo.Context) error {
	buffer := new(bytes.Buffer)
	template.Index(buffer)
	return c.HTML(http.StatusOK, buffer.String())
}

func joinRoom(c echo.Context) error {
	buffer := new(bytes.Buffer)
	template.JoinRoom(c.Param("roomId"), buffer)
	return c.HTML(http.StatusOK, buffer.String())
}

func createRoom(c echo.Context) error {
	clientIp := c.RealIP()
	roomId := base64.StdEncoding.EncodeToString([]byte(clientIp + "TMT"))
	return c.Redirect(http.StatusMovedPermanently, "/room/" + roomId)
}

func serveWebSocket(hub *Hub, roomId string, w http.ResponseWriter, r *http.Request) error {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return err
	}

	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), roomId: roomId}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
	return nil
}