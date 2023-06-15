package router

import (
	"github.com/Pan9Hu/api-server_v2/core"
	"github.com/Pan9Hu/api-server_v2/core/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AttachHealthCheckRoute(engine *gin.Engine) {
	cfg := core.GetAppConfig()
	msg := "this service is running on " + cfg.GetAddress() + ":" + strconv.Itoa(cfg.GetPort())
	// 心跳检测
	engine.GET("ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	// 健康检测
	engine.GET("health", func(ctx *gin.Context) {
		response.OkWithMessage(msg, ctx)
	})
}
