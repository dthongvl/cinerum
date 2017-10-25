package controller

import (
	"github.com/labstack/echo"
	"net/http"
	log "github.com/sirupsen/logrus"
	"github.com/dthongvl/cinerum/src/core/global"
	"github.com/dthongvl/cinerum/src/repository"
	"github.com/dthongvl/cinerum/src/repository/model"
)

func Register(c echo.Context) error {
	return c.String(http.StatusOK, "Register")
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	userId := repository.SignIn(username, password)
	if userId != -1 {
		log.Info(username + " login success")
		cookie, err := global.CookieStore.Get(c.Request(), global.SessionName)
		if err != nil {
			log.Error(err)
		} else {
			cookie.Values["user"] = model.UserCookie{
				IsLoggedIn: true,
				Id:         userId,
				Username:   username,
			}
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
		cookie.Values["user"] = model.UserCookie{}
		err := cookie.Save(c.Request(), c.Response())
		if err != nil {
			log.Error(err)
		}
	}
	referer := c.Request().Header.Get("Referer")
	return c.Redirect(http.StatusFound, referer)
}
