package v2

import (
	"fmt"
	"github.com/Pan9Hu/api-server_v2/service"
	"github.com/Pan9Hu/api-server_v2/util"
	"github.com/gin-gonic/gin"
)

type AccountAPI struct {
}

var (
	stringUtil     util.StringUtil
	accountService service.AccountService
)

func (account *AccountAPI) GetAccountByUID(ctx *gin.Context) {

}

func (account *AccountAPI) ListAccount(ctx *gin.Context) {
	name, _ := stringUtil.SmartTrim(ctx.Query("name"))
	group, _ := stringUtil.SmartTrim(ctx.Query("group"))
	phone, _ := stringUtil.SmartTrim(ctx.Query("phone"))
	email, _ := stringUtil.SmartTrim(ctx.Query("email"))
	sex, _ := stringUtil.SmartTrim(ctx.Query("sex"))
	archGroup, _ := stringUtil.SmartTrim(ctx.Query("arch_group"))
	accountPOJO, err := accountService.ListAccount(name, group, phone, email, sex, archGroup)
	fmt.Println(accountPOJO, err)
}

func (account *AccountAPI) CreateAccount(ctx *gin.Context) {

}

func (account *AccountAPI) DeleteAccountByUID(ctx *gin.Context) {

}

func (account *AccountAPI) DeleteAccount(ctx *gin.Context) {

}

func (account *AccountAPI) UpdateAccountByUID(ctx *gin.Context) {

}
