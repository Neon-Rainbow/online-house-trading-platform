package controller

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProfileGet(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		ResponseErrorWithCode(c, codes.GetDBError)
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		ResponseErrorWithCode(c, codes.GetUserIDError)
		return
	}

	userIDUint, ok := userID.(uint)
	if !ok {
		ResponseErrorWithCode(c, codes.UserIDTypeError)
		return
	}
	userProfile, apiError := logic.GetUserProfile(db, userIDUint)
	if apiError != nil {
		ResponseError(c, *apiError)
		return
	}
	ResponseSuccess(c, userProfile)
}
