package logic

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"

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
