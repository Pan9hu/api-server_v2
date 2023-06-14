package core

import (
	"bytes"
	"fmt"
	"github.com/Pan9Hu/api-server_v2/conf"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"runtime"
)

var (
	config string
	bt     bytes.Buffer
)

// Configure
// 使用viper管理配置
// 优先级 命令行参数 > 环境变量 > 系统默认值
func Configure(filePath ...string) *viper.Viper {

	if filePath[0] == "" {
		//判断命令行参数是否为空
		if configEnv := os.Getenv(conf.ConfigEnv); len(configEnv) < 1 {
			// 判断环境变量是否为空
			if systemType := runtime.GOOS; systemType == "windows" {
				// windows系统默认配置文件路径
				userHomeDir, _ := os.UserHomeDir()
				bt.WriteString(userHomeDir)
				bt.WriteString(conf.DefaultWindowsServerConfigFile)
				config = bt.String()
				fmt.Printf("您正在使用windows系统默认值,config的路径为%s\n", config)
			} else if systemType == "linux" {
				// unix系统默认配置文件路径
				config = conf.DefaultUnixConfigFile
				fmt.Printf("您正在使用unix系统默认值,config的路径为%s\n", config)
			} else {
				// darwin系统默认配置文件路径
				config = conf.DefaultOsxConfigFile
				fmt.Printf("您正在使用darwin系统默认值,config的路径为%s\n", config)
			}
		} else {
			// 环境变量
			config = configEnv
			fmt.Printf("您正在使用环境变量,config的路径为%s\n", config)
		}
	} else {
		// 函数传递值
		config = filePath[0]
		fmt.Printf("您正在使用命令行的-c/--config参数传递的值,config的路径为%s\n", config)
	}

	// 使用viper存储配置
	vp := viper.New()
	vp.SetConfigFile(config)
	vp.SetConfigType("properties")

	err := vp.ReadInConfig() // 检验配置是否可读
	if err != nil {
		panic(fmt.Errorf("[Error] fatal error config file: %s \n", err))
	}

	vp.WatchConfig() // 监控配置有无改变

	vp.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Printf("Config file changed: %v", e.Name)
	})

	return vp
}
