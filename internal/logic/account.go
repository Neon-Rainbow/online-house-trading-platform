package logic

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"
	"os"
)

// DeleteAccountHandle 用于处理删除账户的逻辑
// @title DeleteAccountHandle
// @description 处理删除账户的逻辑
// @param userId uint 用户ID
// @return err *model.Error 错误信息
func DeleteAccountHandle(userId uint) *model.Error {
	err := dao.DeleteAccount(userId)
	if err != nil {
		return &model.Error{StatusCode: codes.LoginServerBusy, Message: codes.LoginServerBusy.Message()}
	}

	user, _ := dao.GetUserProfile(userId)
	if user.Avatar.URL != "" {
		_ = os.Remove(user.Avatar.URL)
	}

	return nil
}
