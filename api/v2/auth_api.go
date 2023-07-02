package v2

import (
	"encoding/base64"
	"github.com/Pan9Hu/api-server_v2/core/response"
	"github.com/Pan9Hu/api-server_v2/dto"
	"github.com/Pan9Hu/api-server_v2/service"
	"github.com/gin-gonic/gin"
	"log"
)

type AuthAPI struct {
}

var (
	authService service.AuthService
)

func (auth *AuthAPI) Login(ctx *gin.Context) {
	params := &dto.LoginDTO{}

	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		log.Printf("[Error] get parames failed: %v ", err.Error())
		return
	}

	username := params.Username
	bPassword, decodeErr := base64.StdEncoding.DecodeString(params.Password)
	password := string(bPassword)
	if decodeErr != nil {
		log.Printf("[Error] decode failed: %v ", decodeErr.Error())
		return
	}

	// 生成X-Auth-Token 前端需要将token放置在localStorage中
	accessToken, refreshToken, tokenErr := authService.LoginByUsername(username, password)
	if tokenErr != nil {
		log.Printf("[Login] Error: %v ", tokenErr.Error())
		if tokenErr.Error() == "the account isn't exists" {
			response.FailWithMessage(tokenErr.Error(), ctx)
			return
		} else if tokenErr.Error() == "wrong password" {
			response.FailWithDetailed(tokenErr.Error(), "30000", nil, ctx)
			return
		} else if tokenErr.Error() == "wrong server" {
			response.FailWithDetailed(tokenErr.Error(), "50000", nil, ctx)
			return
		} else {
			response.FailWithDetailed("fetch token failed", "40000", nil, ctx)
			return
		}
	}
	response.OkWithData(
		gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
			"username":      username,
		},
		ctx,
	)
}

func (auth *AuthAPI) Refresh(ctx *gin.Context) {
	// TODO 刷新过期Token API
}

func (auth *AuthAPI) SecurityCode(ctx *gin.Context) {
	// TODO 忘记密码， 重置密码， 向可能的通知类型发送验证码(邮箱, 手机号码）

}
func (auth *AuthAPI) VerifyCode(ctx *gin.Context) {
	// TODO 验证用户输入的验证码，从定时的数据中核对验证码。
}
func (auth *AuthAPI) ResetPassword(ctx *gin.Context) {
	// TODO 忘记密码， 重置密码。
}
