package main

import (
	"github.com/dthongvl/cinerum/src/module/app"
	"github.com/dthongvl/cinerum/src/module/global"
)

func main() {
	global.Init()
	cinerum := app.New()
	cinerum.RegisterRoute()
	cinerum.Start("3000")
}
