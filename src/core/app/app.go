package app

import (
	"github.com/dthongvl/cinerum/src/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

type App struct {
	port   string
	server *echo.Echo
}

func New() *App {
	a := &App{}
	a.Init()
	return a
}

func (app *App) Init() {
	app.server = echo.New()
	app.port = viper.GetString("server.port")

	app.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: viper.GetString("server.logFormat") + "\n",
	}))
	app.server.Use(middleware.Recover())
}

func (app *App) RegisterRoute() {
	app.server.Static("/static", "static")
	app.server.GET("/", controller.Index)
	app.server.POST("/register", controller.Register)
	app.server.POST("/login", controller.Login)
	app.server.GET("/logout", controller.Logout)
	app.server.GET("/room/:roomID", controller.JoinRoom)
	app.server.GET("/room/create", controller.CreateRoom)
	app.server.GET("/ws", controller.ServeWebSocket)
}

func (app *App) Start() {
	app.server.Logger.Fatal(app.server.Start(":" + app.port))
}
