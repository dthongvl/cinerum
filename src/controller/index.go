package controller

import (
	"bytes"
	"github.com/dthongvl/cinerum/template"
	"github.com/labstack/echo"
	"net/http"
)

func Index(c echo.Context) error {
	buffer := new(bytes.Buffer)
	template.Index(buffer)
	return c.HTML(http.StatusOK, buffer.String())
}
