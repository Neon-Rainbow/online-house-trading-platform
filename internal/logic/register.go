package logic

import (
	"fmt"
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterHandle 用于处理用户注册逻辑
func RegisterHandle(db *gorm.DB, req model.RegisterRequest, c *gin.Context) *model.Error {
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

	dst := fmt.Sprintf("./uploads/user/%d/%d%v", user.ID, user.ID, filepath.Ext(req.Avatar.Filename))
	err = c.SaveUploadedFile(req.Avatar, dst)
	if err != nil {
		return &model.Error{StatusCode: codes.RegisterSaveAvatarError}
	}
	err = dao.CreateUserAvatar(db, &model.UserAvatar{UserID: user.ID, URL: dst})
	if err != nil {
		return &model.Error{StatusCode: codes.RegisterSaveAvatarError}
	}
	return nil
}
