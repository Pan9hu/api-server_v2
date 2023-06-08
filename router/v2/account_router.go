package v2

import (
	"github.com/Pan9Hu/api-server_v2/api"
	"github.com/gin-gonic/gin"
)

type AccountRouter struct {
}

func (r *AccountRouter) InitAccountRouter(Router *gin.RouterGroup) {
	v2AccountAPI := api.APIGroupApp.APIS.AccountAPI
	accountRouter := Router.Group("account")
	{
		accountRouter.GET(":uid", v2AccountAPI.GetAccountByUID)       //通过UID查找账户
		accountRouter.GET("", v2AccountAPI.ListAccount)               // 批量查找账户
		accountRouter.POST("", v2AccountAPI.CreateAccount)            //创建账户
		accountRouter.DELETE(":uid", v2AccountAPI.DeleteAccountByUID) //通过UID删除账户
		accountRouter.DELETE("/", v2AccountAPI.DeleteAccount)         // 批量删除账户
		accountRouter.PUT(":uid", v2AccountAPI.UpdateAccountByUID)    //通过UID更新账户

	}
}
