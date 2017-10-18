package main

import (
	"github.com/dthongvl/cinerum/src/core/app"
	"github.com/dthongvl/cinerum/src/core/chat"
	"github.com/dthongvl/cinerum/src/core/config"
)

func main() {
	config.Load("config.yaml")
	myApp := app.New()
	myApp.RegisterRoute()
	chat.MyHub.Run()
	myApp.Start()
}
