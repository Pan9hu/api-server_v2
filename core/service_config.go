package core

import (
	"github.com/Pan9Hu/api-server_v2/util"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"strings"
	"sync"
)

type ServiceConfig struct {
	address                     string
	port                        int
	mode                        string
	databaseUrl                 string
	databaseMaxIdleConnection   int
	databaseMaxOpenConnection   int
	databaseConnectionMaxIdleMs int64
	databaseConnectionMaxLifeMs int64
}

var (
	appConfig  ServiceConfig
	stringUtil util.StringUtil
	readFile   util.PropertyUtil
)

func (sc *ServiceConfig) GetDatabaseMaxIdleConnection() int {
	return sc.databaseMaxIdleConnection
}

func (sc *ServiceConfig) GetDatabaseMaxOpenConnection() int {
	return sc.databaseMaxOpenConnection
}

func (sc *ServiceConfig) GetDatabaseConnectionMaxIdleMs() int64 {
	return sc.databaseConnectionMaxIdleMs
}

func (sc *ServiceConfig) GetDatabaseConnectionMaxLifeMs() int64 {
	return sc.databaseConnectionMaxLifeMs
}

func (sc *ServiceConfig) SetDatabaseMaxIdleConnection(i int) error {
	sc.databaseMaxIdleConnection = i
	return nil
}

func (sc *ServiceConfig) SetDatabaseMaxOpenConnection(i int) error {
	sc.databaseMaxOpenConnection = i
	return nil
}

func (sc *ServiceConfig) SetDatabaseConnectionMaxIdleMs(i int64) error {
	sc.databaseConnectionMaxIdleMs = i
	return nil
}

func (sc *ServiceConfig) SetDatabaseConnectionMaxLifeMs(i int64) error {
	sc.databaseConnectionMaxLifeMs = i
	return nil
}

func (sc *ServiceConfig) GetAddress() string {
	return sc.address
}

func (sc *ServiceConfig) GetPort() int {
	return sc.port
}

func (sc *ServiceConfig) GetMode() string {
	if sc.mode == "debug" {
		return gin.DebugMode
	} else if sc.mode == "release" {
		return gin.ReleaseMode
	} else {
		return gin.TestMode
	}
}

func (sc *ServiceConfig) SetAddress(address string) bool {
	// TODO 检查address是否合法
	sc.address = address
	return false
}

func (sc *ServiceConfig) SetPort(port int) bool {
	sc.port = port
	return false
}

func (sc *ServiceConfig) SetMode(mode string) bool {
	m := strings.Trim(strings.ToLower(mode), "")
	if len(m) == 0 {
		return false
	}

	if m == "debug" {
		sc.mode = gin.DebugMode
	} else if m == "release" {
		sc.mode = gin.ReleaseMode
	} else if m == "test" {
		sc.mode = gin.TestMode
	} else {
		return false
	}

	return true
}

func (sc *ServiceConfig) SetDatabaseUrl(url string) bool {
	// TODO 验证URL是否合法
	sc.databaseUrl = url
	return true
}

func (sc *ServiceConfig) GetDatabaseUrl() string {
	return sc.databaseUrl
}

func (sc *ServiceConfig) LoadFromFile(filepath string) error {
	if filePath, err := stringUtil.SmartTrim(filepath); err != nil {
		return err
	} else {
		file, fileErr := readFile.ReadPropertiesFile(filePath)
		if fileErr != nil {
			return fileErr
		} else {
			maxIdleMs, _ := strconv.Atoi(file["service.database.connection-max-idle-ms"])
			maxLifeMs, _ := strconv.Atoi(file["service.database.connection-max-life-ms"])
			sc.mode = file["service.mode"]
			sc.port, _ = strconv.Atoi(file["service.port"])
			sc.address = file["service.address"]
			sc.databaseUrl = file["service.database.url"]
			sc.databaseMaxIdleConnection, _ = strconv.Atoi(file["service.database.max-idle-connection"])
			sc.databaseMaxOpenConnection, _ = strconv.Atoi(file["service.database.max-open-connection"])
			sc.databaseConnectionMaxIdleMs = int64(maxIdleMs)
			sc.databaseConnectionMaxLifeMs = int64(maxLifeMs)
			return nil
		}
	}
}

func BuildAppConfig(filepath string) {
	var once sync.Once

	once.Do(func() {
		err := appConfig.LoadFromFile(filepath)
		if err != nil {
			log.Fatalln("load config failed: ", err.Error())
		}
	})
}

func GetAppConfig() *ServiceConfig {
	return &appConfig
}
