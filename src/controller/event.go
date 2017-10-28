package controller

import (
	"bytes"
	"net/http"

	"github.com/CloudyKit/jet"
	"github.com/dthongvl/cinerum/src/core/global"
	"github.com/dthongvl/cinerum/src/repository"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func Events(c echo.Context) error {
	t, err := global.View.GetTemplate("events.jet")
	if err != nil {
		log.Println(err)
		return c.String(http.StatusNoContent, "No content")
	}

	user := getSession(c)
	events := repository.GetEvents()

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
