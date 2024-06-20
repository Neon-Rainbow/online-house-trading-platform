package logic

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/jwt"
	"online-house-trading-platform/pkg/model"
	"time"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

// LoginHandle 用于处理用户登录逻辑
func LoginHandle(db *gorm.DB, req model.LoginRequest, c *gin.Context) (*model.LoginResponse, *model.Error) {
	dbUser, err := dao.GetUserByUsername(db, req.Username)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.LoginUserNotExist}
	}

	if EncryptPassword(req.Password) != dbUser.Password {
		return nil, &model.Error{StatusCode: codes.LoginInvalidPassword}
	}

	token, err := jwt.GenerateToken(dbUser.Username, dbUser.ID, codes.User)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.GenerateJWTTokenError}
	}

	loginRecord := &model.LoginRecord{UserId: dbUser.ID, LoginIp: c.ClientIP(), LoginMethod: c.Request.UserAgent(), LoginTime: time.Now()}
	err = dao.CreateLoginRecord(db, loginRecord)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.LoginServerBusy}
	}

	return &model.LoginResponse{
		Token:    token,
		UserID:   dbUser.ID,
		Username: dbUser.Username,
		Role:     dbUser.Role,
	}, nil
}

// AdminLoginHandle 用于处理管理员登录逻辑
func AdminLoginHandle(db *gorm.DB, req model.LoginRequest, c *gin.Context) (*model.LoginResponse, *model.Error) {
	dbUser, err := dao.GetUserByUsername(db, req.Username)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.LoginUserNotExist}
	}

	if EncryptPassword(req.Password) != dbUser.Password {
		return nil, &model.Error{StatusCode: codes.LoginInvalidPassword}
	}

	if dbUser.Role != "admin" {
		return nil, &model.Error{StatusCode: codes.LoginUserNotExist}
	}

	token, err := jwt.GenerateToken(dbUser.Username, dbUser.ID, codes.Admin)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.GenerateJWTTokenError}
	}

	loginRecord := &model.LoginRecord{UserId: dbUser.ID, LoginIp: c.ClientIP(), LoginMethod: c.Request.UserAgent(), LoginTime: time.Now()}
	err = dao.CreateLoginRecord(db, loginRecord)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.LoginServerBusy}
	}

	return &model.LoginResponse{
		Token:    token,
		UserID:   dbUser.ID,
		Username: dbUser.Username,
		Role:     dbUser.Role,
	}, nil
}
