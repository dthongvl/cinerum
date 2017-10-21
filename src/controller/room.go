package controller

import (
	"bytes"
	"encoding/base64"
	"net/http"

	"github.com/dthongvl/cinerum/src/core/chat"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/CloudyKit/jet"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func getSession(c echo.Context) (isLoggedIn bool, username string) {
	usernameCookie, err := c.Cookie("username")
	if err != nil {
		log.Info("User is not authenticated")
		return false, ""
	}
	log.Info("User is logged in")
	return true, usernameCookie.Value
}

func JoinRoom(c echo.Context) error {
	t, err := View.GetTemplate("room.jet")
	if err != nil {
		return c.String(http.StatusNoContent, "No content")
	}
	var w bytes.Buffer
	vars := make(jet.VarMap)
	isLoggedIn, username := getSession(c)
	vars.Set("roomID", c.Param("roomID"))
	vars.Set("username", username)
	vars.Set("isLoggedIn", isLoggedIn)
	if err = t.Execute(&w, vars, nil); err != nil {
		return c.String(http.StatusNoContent, "No content")
	}
	return c.HTML(http.StatusOK, w.String())
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

	isLoggedIn, username := getSession(c)

	client := &chat.Client{Username: username, IsLoggedIn: isLoggedIn, ChatHub: chat.MyHub, Conn: conn, Send: make(chan []byte, 256), RoomID: c.QueryParam("roomID")}
	client.ChatHub.Register(client)

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
	return nil
}
