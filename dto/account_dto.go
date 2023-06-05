package dto

type AccountDTO struct {
	Username   string `form:"username" json:"username" binding:"required"`
	Password   string `form:"password" json:"password" binding:"required"`
	Name       string `form:"name" json:"name" binding:"required"`
	Phone      string `form:"phone" json:"phone" binding:"required"`
	Email      string `form:"email" json:"email" binding:"required"`
	Group      string `form:"group" json:"group" binding:"required"`
	Sex        string `form:"sex" json:"sex" binding:"required"`
	ArchGroup  string `form:"arch_group" json:"arch_group" binding:"required"`
	CreateTime string `form:"create_time" json:"create_time" binding:"required"`
	UpdateTime string `form:"update_time" json:"update_time" binding:"required"`
}
