package cmd

import (
	"github.com/Pan9Hu/api-server_v2/core"
	"github.com/Pan9Hu/api-server_v2/initialize"
	"github.com/Pan9Hu/api-server_v2/model"
	"github.com/Pan9Hu/api-server_v2/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var smartTrim = &util.StringUtil{}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Start the API Server from a configuration file",
	Long:  `Edit the configuration file(api.properties) and use it to start the API Server`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath, trimErr := smartTrim.SmartTrim(args[0])
		if trimErr != nil {
			log.Fatalln("[Error] started Melo API Server error, because: ", trimErr.Error())
		}
		log.Println("[INFO] starting the Melo API Server...")

		// 初始化服务
		core.BuildAppConfig(filePath)
		config := core.GetAppConfig()
		gin.SetMode(config.GetMode())

		// 初始化数据库
		model.BuildDatabaseTemplate()

		// 初始化路由
		appRoute := initialize.BuildRoute()
		err := appRoute.Run(config.GetAddress() + ":" + strconv.Itoa(config.GetPort()))
		if err != nil {
			log.Fatalln("[Error] started Melo API Server error, because: ", err.Error())
			return
		}
	},
}
