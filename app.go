package main

import "github.com/Pan9Hu/api-server_v2/cmd"

const (
	DEFAULT_WINDOWS_SERVER_CONFIG_FILE = "C:\\etc\\melo-cmdb\\api.properties"
	DEFAULT_UNIX_CONFIG_FILE           = "/etc/melo-cmdb/api.properties"
	DEFAULT_OSX_CONFIG_FILE            = "/etc/melo-cmdb/api.properties"
)

func main() {
	// TODO 环境变量

	// TODO --port --address 参数

	// TODO 默认目录
	cmd.Execute()
}
