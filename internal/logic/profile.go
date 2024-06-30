package logic

import (
	"fmt"
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"
	"os"
	"path/filepath"
)

// GetUserProfile 用于获取用户的个人信息
func GetUserProfile(idUint uint) (*model.User, *model.Error) {
	userProfile, err := dao.GetUserProfile(idUint)
	if err != nil {
		return nil, &model.Error{StatusCode: codes.GetUserProfileError}
	}
	return userProfile, nil
}

// ModifyUserProfile 用于修改用户的个人信息
func ModifyUserProfile(m *model.UserReq, idUint uint) *model.Error {
	user, err := dao.GetUserProfile(idUint)
	if err != nil {
		return &model.Error{StatusCode: codes.GetUserProfileError}
	}
	if user.Role != "admin" {
		m.Role = user.Role
	}
	err = dao.ModifyUserProfile(m, idUint)
	if err != nil {
		return &model.Error{StatusCode: codes.ModifyUserProfileError}
	}
	return nil
}

// ModifyUserAvatar 用于修改用户的头像
func ModifyUserAvatar(avatar *model.UserAvatarReq) *model.Error {
	userInfo, err := dao.GetUserProfile(avatar.UserID)
	if err != nil {
		return &model.Error{StatusCode: codes.GetUserProfileError, Message: "获取用户原先头像信息失败"}
	}
	dst := userInfo.Avatar.URL
	_ = os.Remove(dst)
	fileName := generateRandomFileName()
	dst = fmt.Sprintf("./uploads/user/%d/%s%v", avatar.UserID, fileName, filepath.Ext(avatar.Avatar.Filename))

	err = saveUploadedFile(avatar.Avatar, dst)
	if err != nil {
		return &model.Error{StatusCode: codes.ModifyUserProfileError}
	}
	a := &model.UserAvatar{URL: dst, UserID: avatar.UserID}
	err = dao.ModifyUserAvatar(a)
	if err != nil {
		return &model.Error{StatusCode: codes.ModifyUserProfileError, Message: "修改用户头像信息失败"}
	}

	return nil
}
