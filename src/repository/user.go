package repository

import (
	"github.com/dthongvl/cinerum/src/core/global"
	"github.com/dthongvl/cinerum/src/repository/model"
	"fmt"
	"errors"
)

func SignIn(username string, password string) int {
	query := fmt.Sprintf(model.SignInQuery, username, password)
	user := model.User{}
	err := global.Data.SelectOne(&user, query)
	if err != nil || user.Username != username {
		return -1
	}
	return user.Id
}

func GetStreamSetting(id int) (model.User, error) {
	query := fmt.Sprintf(model.GetStreamSettingQuery, id)
	user := model.User{}
	err := global.Data.SelectOne(&user, query)
	if err != nil {
		return user, errors.New("stream setting not found")
	}
	return user, nil
}