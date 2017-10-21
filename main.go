package main

import (
	"github.com/dthongvl/cinerum/src/core/app"
	"github.com/dthongvl/cinerum/src/core/chat"
	"github.com/dthongvl/cinerum/src/controller"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	controller.View.SetDevelopmentMode(true)
	myApp := app.New()
	myApp.RegisterRoute()
	chat.MyHub.Run()
	myApp.Start()
}
