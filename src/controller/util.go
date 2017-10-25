package controller

import (
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
	"net/http"
	"github.com/CloudyKit/jet"
	"bytes"
	"github.com/dthongvl/cinerum/src/core/global"
	"github.com/labstack/echo"
	"github.com/dthongvl/cinerum/src/repository/model"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}
	return string(result)
}

func errorPage(c echo.Context, user *model.UserCookie, message string) error {
	t, err := global.View.GetTemplate("error.jet")
	if err != nil {
		return c.String(http.StatusNoContent, "No content")
	}
	var w bytes.Buffer
	vars := make(jet.VarMap)
	vars.Set("user", user)
	vars.Set("message", message)
	if err = t.Execute(&w, vars, nil); err != nil {
		return c.String(http.StatusNoContent, "No content")
	}
	return c.HTML(http.StatusOK, w.String())
}

func saveSession(c echo.Context, username string) {
	cookie, err := global.CookieStore.Get(c.Request(), global.SessionName)
	if err != nil {
		log.Error(err)
	} else {
		cookie.Values["user"] = model.UserCookie{
			IsLoggedIn: true,
			RoomID:     username,
		}
		err = cookie.Save(c.Request(), c.Response())
		if err != nil {
			log.Error(err)
		}
	}
}

func getSession(c echo.Context) *model.UserCookie {
	cookie, err := global.CookieStore.Get(c.Request(), global.SessionName)
	if err != nil {
		log.Error(err)
		return &model.UserCookie{}
	}
	user, ok := cookie.Values["user"].(*model.UserCookie)
	if !ok {
		return &model.UserCookie{}
	}
	return user
}

func clearSession(c echo.Context) {
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
}

func addFlash(c echo.Context, message string) {
	cookie, err := global.CookieStore.Get(c.Request(), global.SessionName)
	if err != nil {
		log.Error(err)
	}
	cookie.AddFlash(message)
	cookie.Save(c.Request(), c.Response())
}

func getFlash(c echo.Context) string {
	cookie, err := global.CookieStore.Get(c.Request(), global.SessionName)
	if err != nil {
		log.Error(err)
	}
	flashes := cookie.Flashes()
	lastFlash := ""
	if len(flashes) > 0 {
		lastFlash = flashes[len(flashes) - 1].(string)
	}
	cookie.Save(c.Request(), c.Response())
	return lastFlash
}