package api

import (
	"encoding/base64"
	"github.com/Pan9Hu/api-server_v2/dto"
	"github.com/Pan9Hu/api-server_v2/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AuthAPi struct {
}

func (auth *AuthAPi) Login(ctx *gin.Context) {
	params := &dto.LoginDTO{}

	err := ctx.ShouldBind(&params)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	username := params.Username
	bPassword, err := base64.StdEncoding.DecodeString(params.Password)
	password := string(bPassword)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	authService := &service.AuthService{}
	accessToken, refreshToken, err := authService.LoginByUsername(username, password)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"code":    "10000",
		"data": gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
			"username":      username,
		},
	})
}

func (auth *AuthAPi) Refresh(ctx *gin.Context) {

}

func (auth *AuthAPi) SecurityCode(ctx *gin.Context) {

}
func (auth *AuthAPi) VerifyCode(ctx *gin.Context) {

}
func (auth *AuthAPi) ResetPassword(ctx *gin.Context) {

}
