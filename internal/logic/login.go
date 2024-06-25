package logic

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/jwt"
	"online-house-trading-platform/pkg/model"
	"time"
)

// LoginHandle 用于处理用户登录逻辑
func LoginHandle(req model.LoginRequest, loginIP string, loginMethod string) (*model.LoginResponse, *model.Error) {
	dbUser, err := dao.GetUserByUsername(req.Username)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.LoginUserNotExist}
	}

	if EncryptPassword(req.Password) != dbUser.Password {
		return nil, &model.Error{StatusCode: codes.LoginInvalidPassword}
	}

	accessToken, refreshToken, err := jwt.GenerateToken(dbUser.Username, dbUser.ID, codes.User)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.GenerateJWTTokenError}
	}

	loginRecord := &model.LoginRecord{UserId: dbUser.ID, LoginIp: loginIP, LoginMethod: loginMethod, LoginTime: time.Now()}
	err = dao.CreateLoginRecord(loginRecord)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.LoginServerBusy}
	}

	return &model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       dbUser.ID,
		Username:     dbUser.Username,
		Role:         dbUser.Role,
	}, nil
}

// AdminLoginHandle 用于处理管理员登录逻辑
func AdminLoginHandle(req model.LoginRequest, loginIP string, loginMethod string) (*model.LoginResponse, *model.Error) {
	dbUser, err := dao.GetUserByUsername(req.Username)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.LoginUserNotExist}
	}

	if EncryptPassword(req.Password) != dbUser.Password {
		return nil, &model.Error{StatusCode: codes.LoginInvalidPassword}
	}

	if dbUser.Role != "admin" {
		return nil, &model.Error{StatusCode: codes.LoginUserNotExist}
	}

	accessToken, refreshToken, err := jwt.GenerateToken(dbUser.Username, dbUser.ID, codes.Admin)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.GenerateJWTTokenError}
	}

	loginRecord := &model.LoginRecord{UserId: dbUser.ID, LoginIp: loginIP, LoginMethod: loginMethod, LoginTime: time.Now()}
	err = dao.CreateLoginRecord(loginRecord)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.LoginServerBusy}
	}

	return &model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       dbUser.ID,
		Username:     dbUser.Username,
		Role:         dbUser.Role,
	}, nil
}
