package dto

type AccountDTO struct {
	Uid       string `form:"uid" json:"uid" binding:"required"`
	Name      string `form:"name" json:"name" binding:"required"`
	Phone     string `form:"phone" json:"phone" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
	Group     string `form:"group" json:"group" binding:"required"`
	Sex       string `form:"sex" json:"sex" binding:"required"`
	ArchGroup string `form:"arch_group" json:"arch_group" binding:"required"`
}
