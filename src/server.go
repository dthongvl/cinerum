package src

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func StartWebServer(port string) {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, time=${latency_human}\n",
	}))
	e.Use(middleware.Recover())

	hub := newHub()
	go hub.run()

	e.Static("/static", "static")
	e.GET("/", index)
	e.GET("/room/:roomId", joinRoom)
	e.GET("/room/create", createRoom)
	e.GET("/ws", func(c echo.Context) error {
		serveWebSocket(hub, c.QueryParam("roomId"), c.Response(), c.Request())
		return nil
	})

	e.Logger.Fatal(e.Start(":" + port))
}