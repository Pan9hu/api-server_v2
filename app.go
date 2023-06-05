package main

import (
	api_v2 "github.com/Pan9Hu/api-server_v2/api"
	"github.com/Pan9Hu/api-server_v2/core"
	_ "github.com/Pan9Hu/api-server_v2/core"
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

func buildRoute() *gin.Engine {
	r := gin.Default()
	core.AttachHealthCheckRoute(r)

	v2 := r.Group("/v2")
	{

		v2AuthAPI := new(api_v2.AuthAPi)
		v2AuthAPIGroup := v2.Group("/auth")
		{
			v2AuthAPIGroup.POST("login", v2AuthAPI.Login)
			v2AuthAPIGroup.POST("refresh", v2AuthAPI.Refresh)
			v2AuthAPIGroup.POST("security-code", v2AuthAPI.SecurityCode)
			v2AuthAPIGroup.POST("verify-code", v2AuthAPI.VerifyCode)
			v2AuthAPIGroup.POST("reset-password", v2AuthAPI.ResetPassword)
		}
		v2AccountAPI := new(api_v2.AccountAPI)
		v2AccountAPIGroup := v2.Group("/account")
		{
			v2AccountAPIGroup.GET(":uid", v2AccountAPI.GetAccountByUID)
			v2AccountAPIGroup.GET("", v2AccountAPI.ListAccount)
			v2AccountAPIGroup.POST("", v2AccountAPI.CreateAccount)
			v2AccountAPIGroup.DELETE(":uid", v2AccountAPI.DeleteAccountByUID)
			v2AccountAPIGroup.DELETE("/", v2AccountAPI.DeleteAccount)
			v2AccountAPIGroup.PUT(":uid", v2AccountAPI.UpdateAccountByUID)

		}
		v2GroupAPI := new(api_v2.GroupAPI)
		v2GroupAPIGroup := v2.Group("/group")
		{
			v2GroupAPIGroup.GET(":name", v2GroupAPI.GetGroupByName)
			v2GroupAPIGroup.GET("", v2GroupAPI.GetGroup)
			v2GroupAPIGroup.POST("", v2GroupAPI.CreateGroup)
			v2GroupAPIGroup.DELETE(":name", v2GroupAPI.DeleteGroupByName)
			v2GroupAPIGroup.DELETE("", v2GroupAPI.DeleteGroup)
			v2GroupAPIGroup.PUT(":name", v2GroupAPI.UpdateGroupByName)

		}

	}
	return r
}

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

	appRoute := buildRoute()
	err := appRoute.Run(config.GetAddress() + ":" + strconv.Itoa(config.GetPort()))
	if err != nil {
		log.Fatalln("start melo api server error, because: ", err.Error())
		return
	}
}
