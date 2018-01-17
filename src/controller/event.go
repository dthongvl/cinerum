package controller

import (
	"bytes"
	"net/http"

	"github.com/CloudyKit/jet"
	"github.com/dthongvl/cinerum/src/module/global"
	"github.com/dthongvl/cinerum/src/repository"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/dthongvl/cinerum/src/model"
)

func Events(c echo.Context) error {
	t, err := global.View.GetTemplate("events.jet")
	if err != nil {
		log.Println(err)
		return c.String(http.StatusNoContent, "No content")
	}

	user := getSession(c)
	streamingRooms := repository.GetStreamingRooms()

	var events []model.Event
	for _, room := range streamingRooms {
		event := model.Event{
			RoomId: room.RoomId,
			StreamThumbnail: "/preview/" + room.RoomId + ".jpg",
			StreamTitle: room.StreamTitle,
			TotalOnline: global.ChatHub.GetTotalOnline(room.RoomId),
		}
		events = append(events, event)
	}

	vars := make(jet.VarMap)
	vars.Set("user", user)
	vars.Set("events", events)

	var w bytes.Buffer

	if err = t.Execute(&w, vars, nil); err != nil {
		log.Println(err)
		return c.String(http.StatusNoContent, "No content")
	}
	return c.HTML(http.StatusOK, w.String())
}
