package model

import (
	"mime/multipart"
	"online-house-trading-platform/codes"
)

// LoginRequest 用于处理用户登录请求
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse 用于处理用户登录请求
type LoginResponse struct {
	Token    string `json:"token"`
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
}

// RegisterRequest 用于处理用户注册请求
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

// HouseRequest 是从前端接收房源数据的结构体。
type HouseRequest struct {
	Owner       string                  `json:"owner" form:"owner"`             // 房屋所有者名称
	OwnerID     uint                    `json:"owner_id" form:"owner_id"`       // 房屋所有者ID
	Title       string                  `json:"title" form:"title"`             // 房屋标题
	Description string                  `json:"description" form:"description"` // 房屋描述
	Price       float64                 `json:"price" form:"prices"`            // 房屋价格
	Address     string                  `json:"address" form:"address"`         // 房屋地址
	Images      []*multipart.FileHeader `json:"images" form:"images"`           // 房屋相关的图片文件列表
}

// Error 用于处理错误信息
type Error struct {
	StatusCode codes.ResCode
	Message    string
}
