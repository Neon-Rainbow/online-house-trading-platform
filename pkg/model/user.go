package model

import (
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Username     string        `json:"username" gorm:"uniqueIndex;not null;type:varchar(255)" form:"username" `
	Password     string        `json:"password" gorm:"not null;type:varchar(255)" form:"password" `
	Email        string        `json:"email" gorm:"uniqueIndex;not null;type:varchar(255)" form:"email"`
	Role         string        `json:"role" form:"role" gorm:"type:varchar(255);default:'user'"` // 用于判断用户是否为管理员,默认为普通用户
	PhoneNumber  string        `json:"phone_number" form:"phone_number"`
	Sex          string        `json:"sex" form:"sex"`
	Province     string        `json:"province" form:"province" gorm:"type:varchar(255)"`
	City         string        `json:"city" form:"city" gorm:"type:varchar(255)"`
	Identity     string        `json:"identity" form:"identity" gorm:"type:varchar(255)"`
	QQNumber     string        `json:"qq_number" form:"qq_number" gorm:"type:varchar(255)"`
	WechatNumber string        `json:"wechat_number" form:"wechat_number" gorm:"type:varchar(255)"`
	Avatar       UserAvatar    `json:"avatar" form:"avatar" gorm:"foreignKey:UserID"`
	LoginRecord  []LoginRecord `json:"login_record" gorm:"foreignKey:UserId"`
}

// UserAvatar 用户头像
type UserAvatar struct {
	gorm.Model
	UserID uint   `gorm:"not null" json:"user_id" gorm:"uniqueIndex"`
	URL    string `gorm:"not null" json:"url"`
}
