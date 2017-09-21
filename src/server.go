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

	e.Static("/static", "static")
	e.File("/", "view/index.html")
	e.GET("/room/:roomId", joinRoom)
	e.GET("/room/create", createRoom)
	e.GET("/ws", serveWebSocket)

	go handleMessages()
	e.Logger.Fatal(e.Start(":" + port))
}