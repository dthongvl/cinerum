package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/dthongvl/cinerum/src/repository/model"
)

type Database struct {
	connection *gorm.DB
}

func (db *Database) Init() {
	var err error
	db.connection, err = gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic("failed to connect database")
	}
	db.connection.LogMode(true)
}

func (db *Database) Migrate() {
	db.connection.AutoMigrate(&model.User{})
	user1 := &model.User{
		RoomId: "dthongvl",
		Password: "123456",
		IsDisplay: 1,
		IsPrivate: 0,
		LiveAt: 1,
		StreamTitle: "Untitled",
		StreamKey: "live_dthongvl_zxcqwertyuiop",
	}
	db.connection.Create(user1)

	user2 := &model.User{
		RoomId: "hieuminh",
		Password: "123456",
		IsDisplay: 0,
		IsPrivate: 1,
		LiveAt: 0,
		StreamTitle: "Untitled",
		StreamKey: "live_hieuminh_zxcqwevsvsvsfv",
	}
	db.connection.Create(user2)
}

func (db *Database) GetInstance() (*gorm.DB) {
	return db.connection
}