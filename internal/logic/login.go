package logic

import (
	"context"
	"fmt"
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

	// 使用 context 和 goroutine 并行处理 IP 信息查询和登录记录写入
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resultCh := make(chan error, 1)

	go func() {
		loginIPAddress, loginOperator := GetIPInformationWithContext(ctx, loginIP)
		loginRecord := &model.LoginRecord{
			UserId:      dbUser.ID,
			LoginIp:     loginIP,
			LoginMethod: loginMethod,
			LoginTime:   time.Now(),
			Address:     loginIPAddress,
			Operator:    loginOperator,
		}
		err := dao.CreateLoginRecord(loginRecord)
		resultCh <- err
	}()

	select {
	case err := <-resultCh:
		if err != nil {
			return nil, &model.Error{StatusCode: codes.LoginServerBusy}
		}
	case <-ctx.Done():
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Warning: IP information lookup and login record creation timed out")
		}
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

	// 使用 context 和 goroutine 并行处理 IP 信息查询和登录记录写入
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resultCh := make(chan error, 1)

	go func() {
		loginIPAddress, loginOperator := GetIPInformationWithContext(ctx, loginIP)
		loginRecord := &model.LoginRecord{
			UserId:      dbUser.ID,
			LoginIp:     loginIP,
			LoginMethod: loginMethod,
			LoginTime:   time.Now(),
			Address:     loginIPAddress,
			Operator:    loginOperator,
		}
		err := dao.CreateLoginRecord(loginRecord)
		resultCh <- err
	}()

	select {
	case err := <-resultCh:
		if err != nil {
			return nil, &model.Error{StatusCode: codes.LoginServerBusy}
		}
	case <-ctx.Done():
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Warning: IP information lookup and login record creation timed out")
		}
	}

	return &model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       dbUser.ID,
		Username:     dbUser.Username,
		Role:         dbUser.Role,
	}, nil
}
