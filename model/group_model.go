package model

import "time"

type Group struct {
	GroupName  string    `gorm:"column:group_name; type:varchar(255); primaryKey; not null; comment:组名"`
	Usage      string    `gorm:"column:usage; type:varchar(255);unique; not null; comment:用途"`
	CreateTime time.Time `gorm:"column:create_time; type:datetime; autoCreateTime:milli; not null; comment:创建时间"`
	UpdateTime time.Time `gorm:"column:update_time; type:datetime; autoUpdateTime:milli; not null; comment:更新时间"`
	IsDelete   bool      `gorm:"column:is_delete; type:boolean; not null; comment:是否删除"`
}

func (Group) TableName() string {
	return "t_group"
}
