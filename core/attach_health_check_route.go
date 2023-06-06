package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AttachHealthCheckRoute(engine *gin.Engine) {
	config := GetAppConfig()

	// 心跳检测
	engine.GET("ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")

	})

	// 健康检测
	engine.GET("health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    "10000",
			"message": "this service is running on " + config.GetAddress() + ":" + strconv.Itoa(config.GetPort()),
			"data":    nil,
		})
	})
}
