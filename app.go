package main

import (
	"github.com/Pan9Hu/api-server_v2/core"
	"github.com/Pan9Hu/api-server_v2/initialize"
	"github.com/Pan9Hu/api-server_v2/model"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

const (
	DEFAULT_WINDOWS_SERVER_CONFIG_FILE = "C:\\etc\\melo-cmdb\\api.properties"
	DEFAULT_UNIX_CONFIG_FILE           = "/etc/melo-cmdb/api.properties"
	DEFAULT_OSX_CONFIG_FILE            = "/etc/melo-cmdb/api.properties"
)

func main() {
	// TODO 环境变量

	// TODO --port --address 参数

	// TODO --config 参数

	// TODO 当前目录

	// TODO 默认目录

	core.BuildAppConfig("C:\\Users\\Pan9Hu\\go\\src\\github.com\\Pan9Hu\\api-server_v2\\conf\\api.properties")
	config := core.GetAppConfig()
	model.BuildDatabaseTemplate()
	gin.SetMode(config.GetMode())

	appRoute := initialize.BuildRoute()
	err := appRoute.Run(config.GetAddress() + ":" + strconv.Itoa(config.GetPort()))
	if err != nil {
		log.Fatalln("start melo api server error, because: ", err.Error())
		return
	}
}
