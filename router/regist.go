package router

import routerV2 "github.com/Pan9Hu/api-server_v2/router/v2"

type RouterGroup struct {
	Routers routerV2.Routers
}

var RouterGroupApp = new(RouterGroup)
