package initialize

import (
	"github.com/Pan9Hu/api-server_v2/core"
	"github.com/Pan9Hu/api-server_v2/model"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func InitService() {
	log.Println("[INFO] starting the Melo API Server...")

	// 初始化服务
	core.BuildAppConfig()
	config := core.GetAppConfig()
	gin.SetMode(config.GetMode())

	// 初始化数据库
	model.BuildDatabaseTemplate()

	// 初始化路由
	appRoute := BuildRoute()
	err := appRoute.Run(config.GetAddress() + ":" + strconv.Itoa(config.GetPort()))
	if err != nil {
		log.Fatalln("[Error] started Melo API Server error, because: ", err.Error())
		return
	}
}
