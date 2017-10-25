package main

import (
	"github.com/dthongvl/cinerum/src/core/app"
	"github.com/dthongvl/cinerum/src/core/global"
)

func main() {
	global.Init()
	cinerum := app.New()
	cinerum.RegisterRoute()
	cinerum.Start()
}
