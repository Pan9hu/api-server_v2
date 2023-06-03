package api

import (
	"encoding/base64"
	"fmt"
	"github.com/Pan9Hu/api-server_v2/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type AuthAPi struct {
}

type LoginRequestBodyParams struct {
	Content string `json:"content" binding:"required"`
}

func (auth *AuthAPi) Login(ctx *gin.Context) {
	params := &LoginRequestBodyParams{}

	err := ctx.ShouldBind(&params)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(params.Content)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	fmt.Println(string(decoded))
	decodedContent := strings.Split(string(decoded), ";")
	username := strings.Split(decodedContent[0], "username=")[1]
	password := strings.Split(decodedContent[1], "password=")[1]
	fmt.Println(username + "\n" + password)

	authService := &service.AuthService{}
	accessToken, refreshToken, err := authService.LoginByUsername(username, password)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"code":    "10000",
		"data": gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
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
