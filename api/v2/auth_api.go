package v2

import (
	"encoding/base64"
	"github.com/Pan9Hu/api-server_v2/dto"
	"github.com/Pan9Hu/api-server_v2/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AuthAPI struct {
}

func (auth *AuthAPI) Login(ctx *gin.Context) {
	params := &dto.LoginDTO{}

	err := ctx.ShouldBind(&params)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	username := params.Username
	bPassword, decodeErr := base64.StdEncoding.DecodeString(params.Password)
	password := string(bPassword)
	if decodeErr != nil {
		log.Printf(decodeErr.Error())
		return
	}

	authService := &service.AuthService{}
	accessToken, refreshToken, tokenErr := authService.LoginByUsername(username, password)
	if tokenErr != nil {
		log.Printf("[Login] Error: %v ", tokenErr.Error())
		if tokenErr.Error() == "the account isn't exists" {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "用户不存在",
				"code":    "20000",
				"data": gin.H{
					"access_token":  accessToken,
					"refresh_token": refreshToken,
					"username":      username,
				},
			})
		} else if tokenErr.Error() == "wrong password" {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "密码错误",
				"code":    "30000",
				"data": gin.H{
					"access_token":  accessToken,
					"refresh_token": refreshToken,
					"username":      username,
				},
			})
		} else if tokenErr.Error() == "wrong server" {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "服务不可达",
				"code":    "50000",
				"data": gin.H{
					"access_token":  accessToken,
					"refresh_token": refreshToken,
					"username":      username,
				},
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "获取Token失败",
				"code":    "40000",
				"data": gin.H{
					"access_token":  accessToken,
					"refresh_token": refreshToken,
					"username":      username,
				},
			})
		}
	} else {
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
