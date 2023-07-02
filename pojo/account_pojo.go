package pojo

import "time"

type AccountPOJO struct {
	Username   string
	Name       string
	Password   string
	Phone      string
	Email      string
	Group      string
	Sex        string
	ArchGroup  string
	CreateTime time.Time
	UpdateTime time.Time
	IsDelete   bool
}

var account AccountPOJO

// ID =================

func (ac *AccountPOJO) GetUsername() string {
	return ac.Username
}

func (ac *AccountPOJO) SetUsername(s string) error {
	ac.Username = s
	return nil
}

// Name =================

func (ac *AccountPOJO) GetName() string {
	return ac.Name
}

func (ac *AccountPOJO) SetName(s string) error {
	ac.Name = s
	return nil
}

// Password =================

func (ac *AccountPOJO) GetPassword() string {
	return ac.Password
}

func (ac *AccountPOJO) SetPassword(s string) error {
	ac.Password = s
	return nil
}

// Phone =================

func (ac *AccountPOJO) GetPhone() string {
	return ac.Phone
}

func (ac *AccountPOJO) SetPhone(s string) error {
	ac.Phone = s
	return nil
}

// Email =================

func (ac *AccountPOJO) GetEmail() string {
	return ac.Email
}

func (ac *AccountPOJO) SetEmail(s string) error {
	ac.Email = s
	return nil
}

// Group =================

func (ac *AccountPOJO) GetGroup() string {
	return ac.Group
}

func (ac *AccountPOJO) SetGroup(s string) error {
	ac.Group = s
	return nil
}

// Sex ================

func (ac *AccountPOJO) GetSex() string {
	return ac.Sex
}

func (ac *AccountPOJO) SetSex(s string) error {
	ac.Sex = s
	return nil
}

// ArchGroup ================

func (ac *AccountPOJO) GetArchGroup() string {
	return ac.ArchGroup
}

func (ac *AccountPOJO) SetArchGroup(s string) error {
	ac.ArchGroup = s
	return nil
}

// CreateTime ================

func (ac *AccountPOJO) GetCreateTime() time.Time {
	return ac.CreateTime
}

func (ac *AccountPOJO) SetCreateTime(t time.Time) error {
	ac.CreateTime = t
	return nil
}

// UpdateTime ================

func (ac *AccountPOJO) GetUpdateTime() time.Time {
	return ac.UpdateTime
}

func (ac *AccountPOJO) SetUpdateTime(t time.Time) error {
	ac.UpdateTime = t
	return nil
}

// IsDelete ================

func (ac *AccountPOJO) GetIsDelete() bool {
	return ac.IsDelete
}

func (ac *AccountPOJO) SetIsDelete(b bool) error {
	ac.IsDelete = b
	return nil
}

// =========================

func GetAccount() *AccountPOJO {
	return &account
}
