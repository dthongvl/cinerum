package model

type User struct {
	Id int `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	LiveAt string `db:"live_at"`
	StreamTitle string `db:"stream_title"`
	StreamKey string `db:"stream_key"`
}

var SignInQuery = "SELECT username FROM users WHERE username='%s' AND password='%s'"
//var GetStreamSettingQuery = "SELECT stream_title, stream_key FROM users WHERE username='%s'"