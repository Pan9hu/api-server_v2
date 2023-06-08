package v2

import (
	"github.com/Pan9Hu/api-server_v2/api"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
}

func (r *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {
	v2AuthAPI := api.APIGroupApp.APIS.AuthAPI
	authRouter := Router.Group("auth")
	{
		authRouter.POST("login", v2AuthAPI.Login)                  // 用户登录
		authRouter.POST("refresh", v2AuthAPI.Refresh)              // 刷新Token
		authRouter.POST("security-code", v2AuthAPI.SecurityCode)   // 发送验证码
		authRouter.POST("verify-code", v2AuthAPI.VerifyCode)       // 核对验证码
		authRouter.POST("reset-password", v2AuthAPI.ResetPassword) // 重置密码
	}
}
