package controller

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CollectPost 用于处理用户收藏房屋的Post请求
// @Summary 收藏房屋
// @Description 用户收藏房屋
// @Tags 收藏
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param house_id body uint true "房屋ID"
// @Success 200 {object} controller.ResponseData "收藏成功"
// @Failure 400 {object} controller.ResponseData "预约失败,具体原因查看json中的message字段和code字段"
// @Router /collect [post]
func CollectPost(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		ResponseErrorWithCode(c, codes.GetDBError)
	}

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
	apiError := logic.CollectHandle(db, &favourite, userIDUint)
	if apiError != nil {
		ResponseError(c, *apiError)
	}
	ResponseSuccess(c, nil)
}
