package global

import (
	"github.com/CloudyKit/jet"
	"github.com/gorilla/sessions"
	"github.com/dthongvl/cinerum/src/core/chat"
)

var (
	SessionName = "auth"
	CookieStore = sessions.NewCookieStore([]byte("something-very-secret"))
	View = jet.NewHTMLSet("./template")
	ChatHub = chat.NewHub()
	StreamURL = "rtmp://localhost:1935/app/"
)
