package global

import (
	"encoding/gob"

	"github.com/CloudyKit/jet"
	"github.com/dthongvl/cinerum/src/module/chat"
	"github.com/dthongvl/cinerum/src/repository/model"
	"github.com/gorilla/sessions"
	log "github.com/sirupsen/logrus"
	"github.com/dthongvl/cinerum/src/module/database"
)

var (
	SessionName = "auth"
	CookieStore = sessions.NewCookieStore([]byte("something-very-secret"))
	View        = jet.NewHTMLSet("./template")
	ChatHub     = chat.NewHub()
	StreamURL   = "rtmp://localhost:1935/app/"
	Database    = database.Database{}
)

func Init() {
	log.SetLevel(log.DebugLevel)
	View.SetDevelopmentMode(true)
	ChatHub.Run()
	Database.Init()
	Database.Migrate()
	gob.Register(&model.UserCookie{})
}