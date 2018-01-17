package controller

import (
	"net/http"
	"time"

	"github.com/dthongvl/cinerum/src/repository"
	"github.com/labstack/echo"
	"os/exec"
	"os"
)

func OnPublish(c echo.Context) error {
	streamKey := c.FormValue("name")
	roomID := repository.CheckStreamKey(streamKey)

	if roomID == "" {
		return c.String(http.StatusForbidden, "Access Denied")
	}
	repository.UpdateLiveAt(streamKey, time.Now().Unix())
	go func() {
		for {
			time.Sleep(30 * time.Second)
			room := repository.FindUser(roomID)
			if room.LiveAt == 0 {
				os.Remove("preview/" + roomID + ".jpg")
				break;
			}
			cmd := exec.Command("ffmpeg", "-i", "rtmp://127.0.0.1:1935/hls-live/" + roomID, "-vframes", "1", "-vf", "scale=280:190", "-y", "preview/" + roomID + ".jpg")
			cmd.Run()
		}
	}()
	return c.Redirect(http.StatusFound, roomID)
}

func OnPublishDone(c echo.Context) error {
	streamKey := c.FormValue("name")
	repository.UpdateLiveAt(streamKey, 0)
	return c.String(http.StatusOK, streamKey)
}
