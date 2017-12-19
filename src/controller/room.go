package controller

import (
	"bytes"
	"net/http"

	"strings"

	"github.com/CloudyKit/jet"
	"github.com/dthongvl/cinerum/src/module/chat"
	"github.com/dthongvl/cinerum/src/module/global"
	"github.com/dthongvl/cinerum/src/repository"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func JoinRoom(c echo.Context) error {
	user := getSession(c)
	roomID := c.Param("roomID")
	streamInfo := repository.FindUser(roomID)
	if streamInfo.RoomId == "" {
		log.Println("room not found")
		return errorPage(c, user, "room not found")
	}
	t, err := global.View.GetTemplate("room.jet")
	if err != nil {
		return c.String(http.StatusNoContent, "No content")
	}
	var w bytes.Buffer
	vars := make(jet.VarMap)
	vars.Set("roomID", roomID)
	vars.Set("roomTitle", streamInfo.StreamTitle)
	vars.Set("liveAt", streamInfo.LiveAt)
	vars.Set("user", user)
	if err = t.Execute(&w, vars, nil); err != nil {
		return c.String(http.StatusNoContent, "No content")
	}
	return c.HTML(http.StatusOK, w.String())
}

func RoomSetting(c echo.Context) error {
	user := getSession(c)
	if user.RoomID != c.Param("roomID") {
		log.Println("access another room setting denied")
		return errorPage(c, user, "you do not have permission to access")
	}

	t, err := global.View.GetTemplate("settings.jet")
	if err != nil {
		return c.String(http.StatusNoContent, "No content")
	}
	var w bytes.Buffer
	streamSetting := repository.FindUser(user.RoomID)
	if streamSetting.RoomId == "" {
		referer := c.Request().Header.Get("Referer")
		return c.Redirect(http.StatusMovedPermanently, referer)
	}
	message := getFlash(c)
	errorMessage := ""
	if !strings.Contains(message, "success") {
		errorMessage = message
		message = ""
	}
	vars := make(jet.VarMap)
	vars.Set("user", user)
	vars.Set("streamURL", global.StreamURL)
	vars.Set("settings", streamSetting)
	vars.Set("message", message)
	vars.Set("errorMessage", errorMessage)
	if err = t.Execute(&w, vars, nil); err != nil {
		return c.String(http.StatusNoContent, "No content")
	}
	return c.HTML(http.StatusOK, w.String())
}

func UpdateRoomSetting(c echo.Context) error {
	user := getSession(c)
	roomID := c.Param("roomID")

	if user.RoomID != roomID {
		log.Println("change another room setting denied")
		return errorPage(c, user, "you do not have permission to change")
	}
	if c.FormValue("renewStreamKey") != "" {
		streamInfo := repository.FindUser(roomID)
		if streamInfo.RoomId == "" {
			log.Println("room not found")
			return errorPage(c, user, "room not found")
		}
		if streamInfo.LiveAt != 0 {
			addFlash(c, "cannot renew stream key while streaming")
		} else {
			newStreamKey := "live_" + roomID + "_" + randomString(12)
			repository.UpdateStreamKey(user.RoomID, newStreamKey)
			addFlash(c, "reset successfully")
		}
	} else if c.FormValue("save") != "" {
		isDisplay := 0
		if c.FormValue("isDisplay") == "on" {
			isDisplay = 1
		}
		isPrivate := 0
		if c.FormValue("isPrivate") == "on" {
			isPrivate = 1
		}
		streamTitle := c.FormValue("streamTitle")
		if streamTitle == "" {
			streamTitle = "Untitled"
		}
		repository.UpdateStreamSetting(roomID, isDisplay, isPrivate, streamTitle)
		addFlash(c, "update successfully")
	}

	redirect := "/" + roomID + "/settings"
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
		Username:   user.RoomID,
		IsLoggedIn: user.IsLoggedIn,
		ChatHub:    global.ChatHub,
		Conn:       conn,
		Send:       make(chan []byte, 256),
		RoomID:     c.Param("roomID")}
	client.ChatHub.Register(client)

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
	return nil
}
