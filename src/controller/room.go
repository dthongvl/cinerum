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
	"github.com/dthongvl/cinerum/src/repository/model"
	"github.com/dthongvl/cinerum/src/repository"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func getSession(c echo.Context) *model.UserCookie {
	cookie, err := global.CookieStore.Get(c.Request(), global.SessionName)
	if err != nil {
		log.Error(err)
	}
	user, ok := cookie.Values["user"].(*model.UserCookie)
	if !ok {
		return &model.UserCookie{}
	}
	return user
}

func JoinRoom(c echo.Context) error {
	t, err := global.View.GetTemplate("room.jet")
	if err != nil {
		return c.String(http.StatusNoContent, "No content")
	}
	var w bytes.Buffer
	vars := make(jet.VarMap)
	user := getSession(c)
	vars.Set("roomID", c.Param("roomID"))
	vars.Set("roomTitle", "Room cua " + c.Param("roomID"))
	vars.Set("user", user)
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
	user := getSession(c)
	streamSetting, err := repository.GetStreamSetting(user.Id)
	if err != nil {
		referer := c.Request().Header.Get("Referer")
		return c.Redirect(http.StatusMovedPermanently, referer)
	}
	vars := make(jet.VarMap)
	vars.Set("user", user)
	vars.Set("streamURL", global.StreamURL)
	vars.Set("settings", streamSetting)
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

	user := getSession(c)

	client := &chat.Client{
		Username: user.Username,
		IsLoggedIn: user.IsLoggedIn,
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
