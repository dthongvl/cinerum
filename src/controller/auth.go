package controller

import (
	"net/http"

	"github.com/dthongvl/cinerum/src/repository"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func Register(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	email := c.FormValue("email")
	if !repository.IsUserExist(username, email) {
		repository.Register(username, password, email)
		saveSession(c, username)
	}
	return c.Redirect(http.StatusMovedPermanently, "/")
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
	return c.Redirect(http.StatusMovedPermanently, referer)
}

func Logout(c echo.Context) error {
	clearSession(c)
	referer := c.Request().Header.Get("Referer")
	return c.Redirect(http.StatusMovedPermanently, referer)
}
