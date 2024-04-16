package logic

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/jwt"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
)

// LoginHandle 用于处理用户登录逻辑
func LoginHandle(c *gin.Context, req model.LoginRequest) (*model.LoginResponse, *model.Error) {
	db, err := dao.GetDB(c)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.GetDBError}
	}

	dbUser, err := dao.GetUserByUsername(db, req.Username)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.LoginUserNotExist}
	}

	if EncryptPassword(req.Password) != dbUser.Password {
		return nil, &model.Error{StatusCode: codes.LoginInvalidPassword}
	}

	token, err := jwt.GenerateToken(dbUser.Username, dbUser.ID)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.GenerateJWTTokenError}
	}

	return &model.LoginResponse{
		Token:    token,
		UserID:   dbUser.ID,
		Username: dbUser.Username,
	}, nil
}
