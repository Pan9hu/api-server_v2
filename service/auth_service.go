package service

import (
	"errors"
	"github.com/Pan9Hu/api-server_v2/model"
	"log"
)

type AuthService struct {
}

func (as *AuthService) LoginByUsername(username, password string) (string, string, error) {
	var account *model.Account
	accountPojo := account.AuthUser(username)
	if accountPojo.IsDelete == false {
		if accountPojo.GetUsername() == username {
			if accountPojo.GetPassword() == password {
				if accessToken, accessErr := JWTUtil.GenerateToken(username, "login", true); accessErr == nil {
					if refreshToke, refreshErr := JWTUtil.GenerateToken(username, "login", false); refreshErr == nil {
						// 返回获取的长、短Token
						return accessToken, refreshToke, nil
					} else if refreshErr != nil {
						log.Println("[Info] fetch refresh token failed, because:", refreshErr)
						return "", "", refreshErr
					}
				} else {
					log.Println("[Info] fetch access token failed, because:", accessErr)
					return "", "", accessErr
				}
			} else {
				return "", "", errors.New("wrong password")
			}
		} else {
			return "", "", errors.New("the account isn't exists")
		}
	} else {
		return "", "", errors.New("the account isn't exists")
	}
	return "", "", errors.New("wrong server")
}
