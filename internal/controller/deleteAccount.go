package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"
	"strconv"
)

func DeleteAccount(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		zap.L().Error("RegisterPost: c.MustGet(\"db\").(*gorm.DB) failed",
			zap.String("错误码", strconv.FormatInt(int64(codes.GetDBError), 10)),
		)
		ResponseErrorWithCode(c, codes.GetDBError)
		return
	}

	userId := c.Param("user_id")
	userIdInt, _ := strconv.Atoi(userId)

	apiError := logic.DeleteAccountHandle(db, uint(userIdInt))
	if apiError != nil {
		zap.L().Error("RegisterPost: logic.RegisterHandle failed",
			zap.Int("错误码", apiError.StatusCode.Int()),
		)
		ResponseErrorWithCode(c, apiError.StatusCode)
		return
	}
}
