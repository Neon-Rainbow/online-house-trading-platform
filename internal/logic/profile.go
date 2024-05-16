package logic

import (
	"fmt"
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetUserProfile 用于获取用户的个人信息
func GetUserProfile(db *gorm.DB, idUint uint) (*model.User, *model.Error) {
	userProfile, err := dao.GetUserProfile(db, idUint)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.GetUserProfileError}
	}
	return userProfile, nil
}

func ModifyUserProfile(db *gorm.DB, m *model.UserReq, idUint uint) *model.Error {
	err := dao.ModifyUserProfile(db, m, idUint)
	if err != nil {
		return &model.Error{StatusCode: codes.ModifyUserProfileError}
	}
	return nil
}

func ModifyUserAvatar(db *gorm.DB, avatar *model.UserAvatarReq, c *gin.Context) *model.Error {
	userInfo, err := dao.GetUserProfile(db, avatar.UserID)
	if err != nil {
		return &model.Error{StatusCode: codes.GetUserProfileError, Message: "获取用户原先头像信息失败"}
	}
	dst := userInfo.Avatar.URL
	err = os.Remove(dst)
	if err != nil {
		return &model.Error{StatusCode: codes.DeleteUserAvatarError}
	}
	dst = fmt.Sprintf("./uploads/user/%d/%d%v", avatar.UserID, avatar.UserID, filepath.Ext(avatar.Avatar.Filename))

	err = c.SaveUploadedFile(avatar.Avatar, dst)
	if err != nil {
		return &model.Error{StatusCode: codes.ModifyUserProfileError}
	}

	return nil
}
