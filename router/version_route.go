package router

import (
	"github.com/Pan9Hu/api-server_v2/api"
	"github.com/gin-gonic/gin"
)

func VersionRoute(engine *gin.Engine) {
	// API版本
	versionAPI := new(api.VersionAPI)
	version := engine.Group("api")
	{
		version.GET("version", versionAPI.GetVersion)
	}
}
