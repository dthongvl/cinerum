package app

import (
	"github.com/dthongvl/cinerum/src/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type App struct {
	server *echo.Echo
}

func New() *App {
	a := &App{}
	a.Init()
	return a
}

func (app *App) Init() {
	app.server = echo.New()

	app.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, time=${latency_human}\n",
	}))
	app.server.Use(middleware.Recover())
}

func (app *App) RegisterRoute() {
	app.server.Static("/static", "static")
	app.server.Static("/preview", "preview")
	app.server.GET("/", controller.Index)
	app.server.GET("/events", controller.Events)

	app.server.POST("/register", controller.Register)
	app.server.POST("/login", controller.Login)
	app.server.GET("/logout", controller.Logout)

	app.server.GET("/:roomID", controller.JoinRoom)
	app.server.GET("/:roomID/settings", controller.RoomSetting)
	app.server.POST("/:roomID/settings", controller.UpdateRoomSetting)
	app.server.GET("/:roomID/ws", controller.ServeWebSocket)

	app.server.POST("/on_publish", controller.OnPublish)
	app.server.POST("/on_publish_done", controller.OnPublishDone)
}

func (app *App) Start(port string) {
	app.server.Logger.Fatal(app.server.Start(":" + port))
}
