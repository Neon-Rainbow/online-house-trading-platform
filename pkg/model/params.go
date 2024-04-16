package model

import "online-house-trading-platform/internal/controller"

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

type Error struct {
	StatusCode controller.ResCode
	Message    string
}
