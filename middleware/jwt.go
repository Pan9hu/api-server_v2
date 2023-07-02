package middleware

import (
	"github.com/Pan9Hu/api-server_v2/core/response"
	"github.com/Pan9Hu/api-server_v2/service"
	"github.com/gin-gonic/gin"
)

var jwtService service.JWTService

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//jwt鉴权：头部信息X-Auth-Token
		authHeader := ctx.Request.Header.Get("X-Auth-Token")
		if authHeader == "" {
			return
		}
		code, err := jwtService.TokenIsValid(authHeader)
		if err != nil {
			ctx.Abort()
			response.FailWithDetailed(err.Error(), code, nil, ctx)
			return
		}
		ctx.Next()
	}
}
