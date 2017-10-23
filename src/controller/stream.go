package controller

import (
	"github.com/labstack/echo"
	"strings"
	"net/http"
)

func OnPublish(c echo.Context) error {
	streamKey := c.FormValue("name")

	if strings.HasPrefix(streamKey, "key") {
		return c.Redirect(http.StatusFound, streamKey[3:])
	}
	return c.String(http.StatusForbidden, "Access Denied")
}

func OnPublishDone(c echo.Context) error {
	streamKey := c.FormValue("name")
	return c.String(http.StatusOK, streamKey)
}
