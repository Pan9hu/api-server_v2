package api

import (
	"github.com/Pan9Hu/api-server_v2/core/response"
	"github.com/gin-gonic/gin"
)

type VersionAPI struct {
}

func (version *VersionAPI) GetVersion(ctx *gin.Context) {
	response.OkWithData(
		gin.H{"version": "2.0"},
		ctx,
	)
}
