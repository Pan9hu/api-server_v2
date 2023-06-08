package v2

import (
	"github.com/Pan9Hu/api-server_v2/api"
	"github.com/gin-gonic/gin"
)

type GroupRouter struct {
}

func (r *GroupRouter) InitGroupRouter(Router *gin.RouterGroup) {
	v2GroupAPI := api.APIGroupApp.APIS.GroupAPI
	groupRouter := Router.Group("group")
	{
		groupRouter.GET(":name", v2GroupAPI.GetGroupByName)       //通过组名查找组
		groupRouter.GET("", v2GroupAPI.GetGroup)                  //查找组
		groupRouter.POST("", v2GroupAPI.CreateGroup)              //创建组
		groupRouter.DELETE(":name", v2GroupAPI.DeleteGroupByName) //通过组名删除组
		groupRouter.DELETE("", v2GroupAPI.DeleteGroup)            //删除组
		groupRouter.PUT(":name", v2GroupAPI.UpdateGroupByName)    //通过组名更新组

	}
}
