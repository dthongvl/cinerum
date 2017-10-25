package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/CloudyKit/jet"
	log "github.com/sirupsen/logrus"
	"github.com/dthongvl/cinerum/src/core/global"
	"bytes"
	"github.com/dthongvl/cinerum/src/repository"
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