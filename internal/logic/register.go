package logic

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
)

// RegisterHandle 用于处理用户注册逻辑
func RegisterHandle(c *gin.Context, req model.RegisterRequest) *model.Error {
	db, err := dao.GetDB(c)
	if err != nil {
		return &model.Error{StatusCode: codes.GetDBError}
	}

	if req.Username == "" || req.Password == "" || req.Email == "" {
		return &model.Error{StatusCode: codes.RegisterInvalidParam}
	}

	usernameExists, emailExists, err := dao.CheckUserExists(db, req.Username, req.Email)
	if err != nil {
		return &model.Error{StatusCode: codes.CheckUserExistsError}
	}
	if usernameExists {
		return &model.Error{StatusCode: codes.RegisterUsernameExists}
	}
	if emailExists {
		return &model.Error{StatusCode: codes.RegisterEmailExists}
	}
	user := model.User{
		Username: req.Username,
		Password: EncryptPassword(req.Password),
		Email:    req.Email,
		Role:     req.Role,
	}
	err = dao.CreateUser(db, &user)
	if err != nil {
		return &model.Error{StatusCode: codes.RegisterCreateUserError}
	}
	return &model.Error{StatusCode: codes.CodeSuccess, Message: "注册成功"}
}
