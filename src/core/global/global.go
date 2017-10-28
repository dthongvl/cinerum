package global

import (
	"encoding/gob"

	"github.com/CloudyKit/jet"
	"github.com/dthongvl/cinerum/src/core/chat"
	"github.com/dthongvl/cinerum/src/core/database"
	"github.com/dthongvl/cinerum/src/repository/model"
	"github.com/gorilla/sessions"
	log "github.com/sirupsen/logrus"
)

var (
	SessionName = "auth"
	CookieStore = sessions.NewCookieStore([]byte("something-very-secret"))
	View        = jet.NewHTMLSet("./template")
	ChatHub     = chat.NewHub()
	StreamURL   = "rtmp://localhost:1935/app/"
	Data        = database.Database{}
)

func Init() {
	log.SetLevel(log.DebugLevel)
	View.SetDevelopmentMode(true)
	ChatHub.Run()
	Data.Connect()
	Data.Migrate("schema.sql")
	gob.Register(&model.UserCookie{})
}
