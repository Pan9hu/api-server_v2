package initialize

import (
	"github.com/Pan9Hu/api-server_v2/core"
	"github.com/Pan9Hu/api-server_v2/router"
	"github.com/gin-gonic/gin"
)

func BuildRoute() *gin.Engine {
	Router := gin.Default()
	core.AttachHealthCheckRoute(Router) // 健康检查
	routers := router.RouterGroupApp.Routers
	v2 := Router.Group("v2")
	{
		routers.AuthRouter.InitAuthRouter(v2)
		routers.AccountRouter.InitAccountRouter(v2)
		routers.GroupRouter.InitGroupRouter(v2)
	}
	return Router
}
