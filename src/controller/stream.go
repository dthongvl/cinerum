package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/dthongvl/cinerum/src/repository"
	"time"
)

func OnPublish(c echo.Context) error {
	streamKey := c.FormValue("name")
	roomID, err := repository.CheckStreamKey(streamKey)

	if err != nil {
		return c.String(http.StatusForbidden, "Access Denied")
	}
	repository.UpdateLiveAt(streamKey, time.Now().Unix())
	return c.Redirect(http.StatusFound, roomID)
}

func OnPublishDone(c echo.Context) error {
	streamKey := c.FormValue("name")
	repository.UpdateLiveAt(streamKey, 0)
	return c.String(http.StatusOK, streamKey)
}
