package model

import (
	"github.com/Pan9Hu/api-server_v2/core"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"sync"
	"time"
)

var mysqlTemplate *gorm.DB

func BuildDatabaseTemplate() {
	var once sync.Once

	once.Do(func() {
		config := core.GetAppConfig()
		dsn := config.GetDatabaseUrl()

		template, err1 := gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				// 解决查表自动添加复数问题，例如 user -> users
				SingularTable: true,
			},
		})
		if err1 != nil {
			log.Println(err1.Error())
			panic("[MySQL] Connection failed :" + err1.Error())
			return
		}

		err2 := template.AutoMigrate(&SMSCode{}, &Account{}, &Group{})

		if err2 != nil {
			log.Println(err2.Error())
			panic("[MySQL] Migration failed :" + err2.Error())
			return
		}

		sqlDB, err3 := template.DB()
		if err3 != nil {
			log.Println(err3.Error())
			panic("[MySQL] Configuration failed :" + err2.Error())
			return
		}

		sqlDB.SetMaxIdleConns(config.GetDatabaseMaxIdleConnection())
		sqlDB.SetMaxOpenConns(config.GetDatabaseMaxOpenConnection())
		sqlDB.SetConnMaxIdleTime(time.Millisecond * time.Duration(config.GetDatabaseConnectionMaxIdleMs()))
		sqlDB.SetConnMaxLifetime(time.Millisecond * time.Duration(config.GetDatabaseConnectionMaxLifeMs()))

		mysqlTemplate = template
	})
}

func GetConnectionPool() *gorm.DB {
	return mysqlTemplate
}
