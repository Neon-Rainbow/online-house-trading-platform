package logic

import (
	"gorm.io/gorm"
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"
	"os"
)

func DeleteAccountHandle(db *gorm.DB, userId uint) *model.Error {
	err := dao.DeleteAccount(db, userId)
	if err != nil {
		return &model.Error{StatusCode: codes.LoginServerBusy, Message: codes.LoginServerBusy.Message()}
	}

	user, _ := dao.GetUserProfile(db, userId)
	if user.Avatar.URL != "" {
		_ = os.Remove(user.Avatar.URL)
	}

	return nil
}
