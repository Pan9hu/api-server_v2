package main

import (
	"github.com/Pan9Hu/api-server_v2/cmd"
	"github.com/Pan9Hu/api-server_v2/initialize"
)

func main() {
	cmd.Execute()
	initialize.InitService()
}
