package controller

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
)

// CollectPost 用于处理用户收藏房屋的Post请求
func CollectPost(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		ResponseErrorWithCode(c, codes.GetUserIDError)
	}
	userIDUint, ok := userID.(uint)
	if !ok {
		ResponseErrorWithCode(c, codes.UserIDTypeError)
	}
	var favourite model.Favourite
	err := c.ShouldBind(&favourite)
	if err != nil {
		ResponseErrorWithCode(c, codes.ReserveInvalidParam)
	}
	apiError := logic.CollectHandle(c, &favourite, userIDUint)
	if apiError != nil {
		ResponseError(c, *apiError)
	}
	ResponseSuccess(c, nil)
}
