package service

import (
	"errors"
	"github.com/Pan9Hu/api-server_v2/model"
	"github.com/Pan9Hu/api-server_v2/util"
	"log"
)

type AuthService struct {
}

func (as *AuthService) LoginByUsername(username, password string) (string, string, error) {
	//TODO 查询数据库是否存在这个用户和密码
	var account *model.Account
	accountPojo := account.AuthUser(username)
	if accountPojo.GetID() == username {
		if accountPojo.GetPassword() == password {
			if accessToken, accessErr := util.GenerateToken(username, "login", true); accessErr == nil {
				if refreshToke, refreshErr := util.GenerateToken(username, "login", false); refreshErr == nil {
					// 返回获取的长、短Token
					return accessToken, refreshToke, nil
				} else {
					log.Println(refreshErr)
				}
			} else {
				log.Println(accessErr)
			}
		} else {
			return "", "", errors.New("wrong password")
		}
	} else {
		return "", "", errors.New("the account isn't exists")
	}
	return "", "", errors.New("wrong server")
}
