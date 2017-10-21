package controller

import (
	"github.com/labstack/echo"
	"net/http"
	log "github.com/sirupsen/logrus"
	"github.com/dthongvl/cinerum/src/core/global"
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
		cookie, err := global.CookieStore.Get(c.Request(), global.SessionName)
		if err != nil {
			log.Error(err)
		} else {
			cookie.Values["username"] = username
			err = cookie.Save(c.Request(), c.Response())
			if err != nil {
				log.Error(err)
			}
		}
	}
	referer := c.Request().Header.Get("Referer")
	return c.Redirect(http.StatusSeeOther, referer)
}

func Logout(c echo.Context) error {
	cookie, err := global.CookieStore.Get(c.Request(), global.SessionName)
	if err != nil {
		log.Error(err)
	} else {
		cookie.Values["username"] = ""
		err := cookie.Save(c.Request(), c.Response())
		if err != nil {
			log.Error(err)
		}
	}
	return c.Redirect(http.StatusFound, "/")
}
