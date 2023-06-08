package model

import (
	"github.com/Pan9Hu/api-server_v2/pojo"
	"time"
)

type Account struct {
	ID         string    `gorm:"column:id; type:varchar(255); primaryKey; comment:工号"`
	Name       string    `gorm:"column:name; type:varchar(255);unique ;not null; comment:姓名"`
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

func (ac *Account) TableName() string {
	return "t_account"
}

func (ac *Account) GetAccountByUID(id string) *pojo.AccountPOJO {
	return &pojo.AccountPOJO{
		ID:         ac.ID,
		Name:       ac.Name,
		Password:   ac.Password,
		Phone:      ac.Phone,
		Email:      ac.Email,
		Group:      ac.Group,
		Sex:        ac.Sex,
		ArchGroup:  ac.ArchGroup,
		CreateTime: ac.CreateTime,
		UpdateTime: ac.UpdateTime,
		IsDelete:   ac.IsDelete,
	}
}

func (ac *Account) AllAccount() []any {
	var accounts []any
	return accounts
}

func (ac *Account) GetAccount() []any {
	// TODO 范围查找
	var accounts []any
	return accounts
}

func (ac *Account) CreateAccount(ID string, Name string, Phone string, Email string, Group string, Sex string,
	ArchGroup string, CreateTime time.Time, UpdateTime time.Time) error {
	return nil
}

func (ac *Account) UpdateAccount(ID string, UpdateDict map[string]any) error {
	return nil
}

func (ac *Account) DeleteAccountByID(id string) error {
	return nil
}

func (ac *Account) DeleteAccount(DeleteList []string) error {
	return nil
}

func (ac *Account) AuthUser(username string) *pojo.AccountPOJO {
	db := GetConnectionPool()
	db.Where("id = ?", username).Find(&ac)
	return &pojo.AccountPOJO{
		ID:       ac.ID,
		Password: ac.Password,
		IsDelete: ac.IsDelete,
	}
}
