package model

import "github.com/jinzhu/gorm"

type UserCookie struct {
	IsLoggedIn bool
	RoomID     string
}

type User struct {
	gorm.Model
	RoomId      string
	Password    string
	Email       string
	IsDisplay   int    `gorm:"default:1"`
	IsPrivate   int    `gorm:"default:0"`
	LiveAt      int64  `gorm:"default:0"`
	StreamTitle string `gorm:"default:'Untitled'"`
	StreamKey   string
}
