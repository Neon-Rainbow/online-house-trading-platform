package logic

import (
	"crypto/md5"
	"encoding/hex"
	"online-house-trading-platform/config"
	"online-house-trading-platform/internal/controller"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/jwt"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
)

// EncryptPassword 用于对密码进行加密,使用了md5加密算法
func EncryptPassword(password string) string {
	// secret 用于存储密码加密的密钥
	var secret = config.AppConfig.PasswordSecret
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

// Login 用于处理用户登录逻辑
func Login(c *gin.Context, req model.LoginRequest) (*model.LoginResponse, *model.Error) {
	db, err := dao.GetDB(c)
	if err != nil {
		return nil, &model.Error{StatusCode: controller.GetDBError}
	}

	dbUser, err := dao.GetUserByUsername(db, req.Username)
	if err != nil {
		return nil, &model.Error{StatusCode: controller.LoginUserNotExist}
	}

	if EncryptPassword(req.Password) != dbUser.Password {
		return nil, &model.Error{StatusCode: controller.LoginInvalidPassword}
	}

	token, err := jwt.GenerateToken(dbUser.Username, dbUser.ID)
	if err != nil {
		return nil, &model.Error{StatusCode: controller.GenerateJWTTokenError}
	}

	return &model.LoginResponse{
		Token:    token,
		UserID:   dbUser.ID,
		Username: dbUser.Username,
	}, nil
}
