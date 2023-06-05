package service

import (
	"fmt"
	"github.com/Pan9Hu/api-server_v2/model"
)

type AuthService struct {
}

func (as *AuthService) LoginByUsername(username, password string) (string, string, error) {
	//TODO 查询数据库是否存在这个用户和密码
	dbAccount := &model.Account{}
	result := dbAccount.AuthUser(username)
	fmt.Println(result)
	return "s", "s", nil
}
