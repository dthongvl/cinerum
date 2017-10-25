package model

type UserCookie struct {
	IsLoggedIn bool
	Id int
	Username string
}

type User struct {
	Id          int    `db:"id"`
	Username    string `db:"username"`
	Password    string `db:"password"`
	IsDisplay   int   `db:"is_display"`
	IsPrivate   int   `db:"is_private"`
	LiveAt      string `db:"live_at"`
	StreamTitle string `db:"stream_title"`
	StreamKey   string `db:"stream_key"`
}

var (
	SignInQuery           = "SELECT id, username FROM users WHERE username='%s' AND password='%s'"
	GetStreamSettingQuery = "SELECT is_display, is_private, stream_title, stream_key FROM users WHERE id=%d"
)
