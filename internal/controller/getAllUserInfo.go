package controller

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetAllUsersInformation 用于处理管理员获取所有用户信息的Get请求
func GetAllUsersInformation(c *gin.Context) {
	includeDeleted := c.Query("include_deleted")

	users, err := logic.GetAllUsers(includeDeleted)
	if err != nil {
		zap.L().Error("GetAllUsersInformation: logic.GetAllUsers failed",
			zap.String("错误码", strconv.FormatInt(int64(codes.GetAllUsersError), 10)),
		)
		ResponseErrorWithCode(c, codes.GetAllUsersError)
		return
	}
	ResponseSuccess(c, users)
	return
}
