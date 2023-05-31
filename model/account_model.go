package model

import "time"

type Account struct {
	ID         string    `gorm:"column:id; type:varchar(255); primaryKey; comment:工号"`
	Username   string    `gorm:"column:username; type:varchar(255)unique ;not null; comment:姓名"`
	Password   string    `gorm:"column:password; type:varchar(255); unique; not null; comment:密码"`
	Phone      string    `gorm:"column:phone; type:varchar(255); unique; not null; comment:手机号码"`
	Email      string    `gorm:"column:email; type:varchar(255); unique; not null; comment:工作邮箱"`
	Group      string    `gorm:"column:group; type:varchar(255); comment:组"`
	Sex        string    `gorm:"column:sex; type:varchar(255); comment:性别"`
	ArchGroup  string    `gorm:"column:arch_group; type:varchar(255); comment:组织架构"`
	CreateTime time.Time `gorm:"column:create_time; type:datetime; autoCreateTime:milli; not null; comment:创建时间"`
	UpdateTime time.Time `gorm:"column:update_time; type:datetime; autoUpdateTime:milli; not null; comment:更新时间"`
	IsDelete   bool      `gorm:"column:is_delete; type:boolean; not null; comment:是否删除"`
}

func (Account) TableName() string {
	return "t_account"
}
