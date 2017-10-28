package repository

import (
	"errors"
	"fmt"

	"github.com/dthongvl/cinerum/src/core/global"
	"github.com/dthongvl/cinerum/src/repository/model"
)

func SignIn(roomID string, password string) (string, error) {
	query := fmt.Sprintf(model.SignInQuery, roomID, password)
	user := model.User{}
	err := global.Data.SelectOne(&user, query)
	return user.RoomID, err
}

func GetStreamSetting(roomID string) (model.User, error) {
	query := fmt.Sprintf(model.GetStreamSettingQuery, roomID)
	user := model.User{}
	err := global.Data.SelectOne(&user, query)
	if err != nil {
		return user, errors.New("stream setting not found")
	}
	return user, nil
}

func GetStreamInfo(roomID string) (model.User, error) {
	query := fmt.Sprintf(model.GetStreamInfoQuery, roomID)
	user := model.User{}
	err := global.Data.SelectOne(&user, query)
	if err != nil {
		return user, errors.New("stream info not found")
	}
	return user, nil
}

func UpdateStreamSetting(roomID string, isDisplay int, isPrivate int, streamTitle string) {
	query := fmt.Sprintf(model.UpdateStreamSettingQuery, isDisplay, isPrivate, streamTitle, roomID)
	global.Data.UpdateDelete(query)
}

func UpdateStreamKey(roomID string, newStreamKey string) {
	query := fmt.Sprintf(model.UpdateStreamKeyQuery, newStreamKey, roomID)
	global.Data.UpdateDelete(query)
}

func CheckStreamKey(streamKey string) (string, error) {
	query := fmt.Sprintf(model.GetStreamKeyQuery, streamKey)
	user := model.User{}
	err := global.Data.SelectOne(&user, query)
	return user.RoomID, err
}

func UpdateLiveAt(roomID string, liveAt int64) {
	query := fmt.Sprintf(model.UpdateLiveAtQuery, liveAt, roomID)
	global.Data.UpdateDelete(query)
}

func GetEvents() []model.User {
	query := model.GetEventsQuery
	user := []model.User{}
	global.Data.Select(&user, query)
	return user
}
