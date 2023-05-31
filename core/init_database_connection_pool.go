package core

import (
	"github.com/Pan9Hu/api-server_v2/model"
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
		config := GetAppConfig()
		dsn := config.GetDatabaseUrl()

		mysqlTemplate, err1 := gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				// 解决查表自动添加复数问题，例如 user -> users
				SingularTable: true,
			},
		})
		if err1 != nil {
			log.Println(err1.Error())
			return
		}

		err2 := mysqlTemplate.AutoMigrate(&model.SMSCode{}, &model.Account{}, &model.Group{})

		if err2 != nil {
			log.Println(err2.Error())
			return
		}

		sqlDB, err3 := mysqlTemplate.DB()
		if err3 != nil {
			log.Println(err3.Error())
			return
		}

		sqlDB.SetMaxIdleConns(config.GetDatabaseMaxIdleConnection())
		sqlDB.SetMaxOpenConns(config.GetDatabaseMaxOpenConnection())
		sqlDB.SetConnMaxIdleTime(time.Millisecond * time.Duration(config.GetDatabaseConnectionMaxIdleMs()))
		sqlDB.SetConnMaxLifetime(time.Millisecond * time.Duration(config.GetDatabaseConnectionMaxLifeMs()))

	})

}

func GetConnectionPool() *gorm.DB {
	return mysqlTemplate
}
