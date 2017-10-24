package controller

import (
	"bytes"
	"net/http"

	"github.com/dthongvl/cinerum/src/core/chat"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/CloudyKit/jet"
	"github.com/dthongvl/cinerum/src/core/global"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func getSession(c echo.Context) (isLoggedIn bool, username string) {
	cookie, err := global.CookieStore.Get(c.Request(), global.SessionName)
	if err != nil {
		log.Error(err)
	}
	username, ok := cookie.Values["username"].(string);
	if !ok || username == "" {
		log.Info("User is not authenticated")
		return false, ""
	}
	log.Info("User is logged in")
	return true, username
}

func JoinRoom(c echo.Context) error {
	t, err := global.View.GetTemplate("room.jet")
	if err != nil {
		return c.String(http.StatusNoContent, "No content")
	}
	var w bytes.Buffer
	vars := make(jet.VarMap)
	isLoggedIn, username := getSession(c)
	vars.Set("roomID", c.Param("roomID"))
	vars.Set("roomTitle", "Room cua " + c.Param("roomID"))
	vars.Set("username", username)
	vars.Set("isLoggedIn", isLoggedIn)
	if err = t.Execute(&w, vars, nil); err != nil {
		return c.String(http.StatusNoContent, "No content")
	}
	return c.HTML(http.StatusOK, w.String())
}

func RoomSetting(c echo.Context) error {
	t, err := global.View.GetTemplate("setting.jet")
	if err != nil {
		return c.String(http.StatusNoContent, "No content")
	}
	var w bytes.Buffer
	vars := make(jet.VarMap)
	isLoggedIn, username := getSession(c)
	vars.Set("isLoggedIn", isLoggedIn)
	vars.Set("username", username)
	vars.Set("streamTitle", "title")
	vars.Set("streamURL", global.StreamURL)
	vars.Set("streamKey", "key" + username)
	if err = t.Execute(&w, vars, nil); err != nil {
		return c.String(http.StatusNoContent, "No content")
	}
	return c.HTML(http.StatusOK, w.String())
}

func UpdateRoomSetting(c echo.Context) error {
	if c.FormValue("resetStreamKey") != "" {

	} else if c.FormValue("save") != "" {

	}
	redirect := "/" + c.Param("roomID") + "/setting"
	return c.Redirect(http.StatusMovedPermanently, redirect)
}

func ServeWebSocket(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("upgrade:", err)
		return err
	}

	isLoggedIn, username := getSession(c)

	client := &chat.Client{
		Username: username,
		IsLoggedIn: isLoggedIn,
		ChatHub: global.ChatHub,
		Conn: conn,
		Send: make(chan []byte, 256),
		RoomID: c.Param("roomID")}
	client.ChatHub.Register(client)

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
	return nil
}
