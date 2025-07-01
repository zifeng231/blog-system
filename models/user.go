package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique" json:"username" binding:"required,min=3,max=20" vmsg:"required:用户名不能为空,min:用户名长度至少3位,max:用户名长度不能超过20位"`
	Password string `json:"password" binding:"required,min=6" vmsg:"required:密码不能为空,min:密码长度至少3位"`
	Email    string `gorm:"unique" json:"email" binding:"required,email" vmsg:"required:邮箱不能为空,email:邮箱格式不正确"`
}
