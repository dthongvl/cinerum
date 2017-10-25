package model

type UserCookie struct {
	IsLoggedIn bool
	RoomID     string
}

type User struct {
	Id          int    `db:"id"`
	RoomID      string `db:"room_id"`
	Password    string `db:"password"`
	IsDisplay   int    `db:"is_display"`
	IsPrivate   int    `db:"is_private"`
	LiveAt      string `db:"live_at"`
	StreamTitle string `db:"stream_title"`
	StreamKey   string `db:"stream_key"`
}

var (
	SignInQuery              = "SELECT room_id FROM users WHERE room_id='%s' AND password='%s'"
	GetStreamSettingQuery    = "SELECT is_display, is_private, stream_title, stream_key FROM users WHERE room_id='%s'"
	UpdateStreamSettingQuery = "UPDATE users SET is_display=%d, is_private=%d, stream_title='%s' WHERE room_id='%s'"
	UpdateStreamKeyQuery     = "UPDATE users SET stream_key='%s' WHERE room_id='%s'"
)
