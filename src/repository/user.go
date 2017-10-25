package repository

import (
	"github.com/dthongvl/cinerum/src/core/global"
	"github.com/dthongvl/cinerum/src/repository/model"
	"fmt"
	"errors"
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

func UpdateStreamSetting(roomID string, isDisplay int, isPrivate int, streamTitle string) {
	query := fmt.Sprintf(model.UpdateStreamSettingQuery, isDisplay, isPrivate, streamTitle, roomID)
	global.Data.UpdateDelete(query)
}

func UpdateStreamKey(roomID string, newStreamKey string) {
	query := fmt.Sprintf(model.UpdateStreamKeyQuery, newStreamKey, roomID)
	global.Data.UpdateDelete(query)
}