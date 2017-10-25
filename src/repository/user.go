package repository

import (
	"github.com/dthongvl/cinerum/src/core/global"
	"github.com/dthongvl/cinerum/src/repository/model"
	"fmt"
)

func SignIn(username string, password string) bool {
	query := fmt.Sprintf(model.SignInQuery, username, password)
	user := model.User{}
	err := global.Data.SelectOne(&user, query)
	if err != nil {
		return false
	}
	return user.Username == username
}

//func GetStreamSetting(username string) model.User {
//	query := fmt.Sprintf(model.GetStreamSettingQuery, username)
//	user := model.User{}
//	err := global.Data.SelectOne(user, query)
//	if err != nil {
//
//	}
//}