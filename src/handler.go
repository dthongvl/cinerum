package src

import (
	"github.com/labstack/echo"
	"net/http"
	"encoding/base64"
)

func joinRoom(c echo.Context) error {
	return c.File("view/room.html")
}

func createRoom(c echo.Context) error {
	clientIp := c.RealIP()
	roomId := base64.StdEncoding.EncodeToString([]byte(clientIp + "TMT"))
	return c.Redirect(http.StatusMovedPermanently, "/room/" + roomId)
}