package global

import (
	"github.com/CloudyKit/jet"
	"github.com/gorilla/sessions"
	"github.com/dthongvl/cinerum/src/core/chat"
	"github.com/dthongvl/cinerum/src/core/database"
	"encoding/gob"
	"github.com/dthongvl/cinerum/src/repository/model"
	log "github.com/sirupsen/logrus"
)

var (
	SessionName = "auth"
	CookieStore = sessions.NewCookieStore([]byte("something-very-secret"))
	View = jet.NewHTMLSet("./template")
	ChatHub = chat.NewHub()
	StreamURL = "rtmp://localhost:1935/app/"
	Data = database.Database{}
)

func Init() {
	log.SetLevel(log.DebugLevel)
	View.SetDevelopmentMode(true)
	ChatHub.Run()
	Data.Connect()
	Data.Migrate("schema.sql")
	gob.Register(&model.UserCookie{})
}