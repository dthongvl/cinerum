package controller

import (
	"net/http"

	"github.com/dthongvl/cinerum/src/repository"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func Register(c echo.Context) error {
	return c.String(http.StatusOK, "Register")
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	roomID := repository.SignIn(username, password)
	if roomID != "" {
		log.Info(username + " login success")
		saveSession(c, roomID)
	}
	referer := c.Request().Header.Get("Referer")
	return c.Redirect(http.StatusSeeOther, referer)
}

func Logout(c echo.Context) error {
	clearSession(c)
	referer := c.Request().Header.Get("Referer")
	return c.Redirect(http.StatusFound, referer)
}
