package controller

import (
	log "github.com/sirupsen/logrus"
	"github.com/labstack/echo"
	"net/http"
	"bytes"
	"github.com/dthongvl/cinerum/src/core/global"
	"github.com/CloudyKit/jet"
)

func Index(c echo.Context) error {
	t, err := global.View.GetTemplate("index.jet")
	if err != nil {
		log.Println(err)
		return c.String(http.StatusNoContent, "No content")
	}
	var w bytes.Buffer
	vars := make(jet.VarMap)
	isLoggedIn, _ := getSession(c)
	vars.Set("isLoggedIn", isLoggedIn)
	if err = t.Execute(&w, vars, nil); err != nil {
		log.Println(err)
		return c.String(http.StatusNoContent, "No content")
	}
	return c.HTML(http.StatusOK, w.String())
}
