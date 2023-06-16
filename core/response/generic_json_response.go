package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GenericJSONResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Data    any    `json:"data"`
}

const (
	SUCCESS = "10000"
	ERROR   = "20000"
)

func Build(message string, code string, data any, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, GenericJSONResponse{
		message,
		code,
		data,
	})
}

func Ok(ctx *gin.Context) {
	Build("完成", SUCCESS, nil, ctx)
}
func OkWithMessage(message string, ctx *gin.Context) {
	Build(message, SUCCESS, nil, ctx)
}

func OkWithData(data any, ctx *gin.Context) {
	Build("完成", SUCCESS, data, ctx)
}

func OkWithDetailed(message string, code string, data any, ctx *gin.Context) {
	Build(message, code, data, ctx)
}

func Fail(ctx *gin.Context) {
	Build("失败", ERROR, nil, ctx)
}

func FailWithMessage(message string, ctx *gin.Context) {
	Build(message, ERROR, nil, ctx)
}

func FailWithDetailed(message string, code string, data any, ctx *gin.Context) {
	Build(message, code, data, ctx)
}
