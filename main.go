package main

import (
	"github.com/dthongvl/cinerum/src/core/app"
	"github.com/dthongvl/cinerum/src/core/global"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	global.View.SetDevelopmentMode(true)
	myApp := app.New()
	myApp.RegisterRoute()
	global.MyHub.Run()
	myApp.Start()
}
