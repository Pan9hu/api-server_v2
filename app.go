package main

import (
	"github.com/Pan9Hu/api-server_v2/cmd"
	"github.com/Pan9Hu/api-server_v2/initialize"
)

func main() {
	// TODO 环境变量

	// TODO --port --address 参数

	// TODO 默认目录
	cmd.Execute()
	initialize.InitService()
}
