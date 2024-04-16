package logic

import (
	"online-house-trading-platform/internal/controller"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
)

// RegisterHandle 用于处理用户注册逻辑
func RegisterHandle(c *gin.Context, req model.RegisterRequest) *model.Error {
	db, err := dao.GetDB(c)
	if err != nil {
		return &model.Error{StatusCode: controller.GetDBError}
	}

	if req.Username == "" || req.Password == "" || req.Email == "" {
		return &model.Error{StatusCode: controller.RegisterInvalidParam}
	}

	usernameExists, emailExists, err := dao.CheckUserExists(db, req.Username, req.Email)
	if err != nil {
		return &model.Error{StatusCode: controller.CheckUserExistsError}
	}
	if usernameExists {
		return &model.Error{StatusCode: controller.RegisterUsernameExists}
	}
	if emailExists {
		return &model.Error{StatusCode: controller.RegisterEmailExists}
	}
	user := model.User{
		Username: req.Username,
		Password: EncryptPassword(req.Password),
		Email:    req.Email,
		Role:     req.Role,
	}
	err = dao.CreateUser(db, &user)
	if err != nil {
		return &model.Error{StatusCode: controller.RegisterCreateUserError}
	}
	return &model.Error{StatusCode: controller.CodeSuccess, Message: "注册成功"}
}
