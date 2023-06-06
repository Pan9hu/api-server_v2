package model

import (
	"time"
)

type SMSCode struct {
	ID          string    `gorm:"column:id; type:varchar(255);primaryKey ; comment:工号"`
	AuthMethod  string    `gorm:"column:auth_method; type:varchar(255); not null; comment:密码"`
	Phone       string    `gorm:"column:phone; type:varchar(255); unique; comment:手机号码"`
	Email       string    `gorm:"column:email; type:varchar(255); unique; comment:工作邮箱"`
	Code        string    `gorm:"column:code; type:varchar(255); unique; not null;comment:工作邮箱"`
	ExpiredTime time.Time `gorm:"column:expired_time; type:datetime; autoUpdateTime:milli; not null; comment:过期时间"`
}

func (sc *SMSCode) TableName() string {
	return "t_sms_code"
}
