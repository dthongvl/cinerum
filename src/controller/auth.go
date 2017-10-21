package controller

import (
	"github.com/labstack/echo"
	"net/http"
	log "github.com/sirupsen/logrus"
)

func Register(c echo.Context) error {
	return c.String(http.StatusOK, "Register")
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	log.Info("NEW LOGIN: " + username + " - " + password)
	if password == "123456" {
		log.Info(username + " LOGIN SUCCESS")
		cookie := new(http.Cookie)
		cookie.Name = "username"
		cookie.Value = username
		c.SetCookie(cookie)
	}
	referer := c.Request().Header.Get("Referer")
	return c.Redirect(http.StatusFound, referer)
}

func Logout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = ""
	cookie.MaxAge = -1
	c.SetCookie(cookie)
	return c.Redirect(http.StatusFound, "/")
}