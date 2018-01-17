package repository

import (
	"github.com/dthongvl/cinerum/src/module/global"
	"github.com/dthongvl/cinerum/src/model"
)

func SignIn(roomID string, password string) (string) {
	var user model.User
	global.Database.GetInstance().
		Where("room_id = ? AND password = ?", roomID, password).
		First(&user)
	return user.RoomId
}

func IsUserExist(roomID string, email string) bool {
	var user model.User
	global.Database.GetInstance().
		Where("room_id = ? OR email = ?", roomID, email).
		First(&user)
	return user.RoomId == roomID || user.Email == email
}

func Register(roomID string, password string, email string) {
	user := model.User{
		RoomId:   roomID,
		Password: password,
		Email:    email,
	}
	global.Database.GetInstance().Create(&user)
}

func FindUser(roomID string) (model.User) {
	var user model.User
	global.Database.GetInstance().
		Where("room_id = ?", roomID).
		First(&user)
	return user
}

func UpdateStreamSetting(roomID string, isDisplay int, isPrivate int, streamTitle string) {
	global.Database.GetInstance().
		Model(&model.User{}).
		Where("room_id = ?", roomID).
		Update(map[string]interface{}{
		"is_display":   isDisplay,
		"is_private":   isPrivate,
		"stream_title": streamTitle,
	})
}

func UpdateStreamKey(roomID string, newStreamKey string) {
	global.Database.GetInstance().
		Model(&model.User{}).
		Where("room_id = ?", roomID).
		Update("stream_key", newStreamKey)
}

func CheckStreamKey(streamKey string) (string) {
	var user model.User
	global.Database.GetInstance().
		Where("stream_key = ?", streamKey).
		First(&user)
	return user.RoomId
}

func UpdateLiveAt(streamKey string, liveAt int64) {
	global.Database.GetInstance().
		Model(&model.User{}).
		Where("stream_key = ?", streamKey).
		Update("live_at", liveAt)
}

func GetStreamingRooms() []model.User {
	var users []model.User
	global.Database.GetInstance().
		Where("live_at > 0 AND is_display = 1").
		Find(&users)
	return users
}
