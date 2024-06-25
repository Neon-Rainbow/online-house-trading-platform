package controller

import (
	"online-house-trading-platform/internal/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func DeleteUserAccountByUserID(c *gin.Context) {
	userId := c.Param("user_id")
	userIdInt, _ := strconv.Atoi(userId)

	apiError := logic.DeleteAccountHandle(uint(userIdInt))
	if apiError != nil {
		zap.L().Error("RegisterPost: logic.RegisterHandle failed",
			zap.Int("错误码", apiError.StatusCode.Int()),
		)
		ResponseErrorWithCode(c, apiError.StatusCode)
		return
	}
	ResponseSuccess(c, nil)
}
