package controller

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"
	"online-house-trading-platform/pkg/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// ProfileGet 用于处理用户获取个人信息的Get请求
// @Summary 获取个人信息
// @Description 获取个人信息
// @Tags 个人信息
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param user_id path string true "用户ID"
// @Success 200 {object} controller.ResponseData "获取成功"
// @Failure 400 {object} controller.ResponseData "获取失败"
// @Router /profile/{user_id} [get]
func ProfileGet(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		zap.L().Error("ProfileGet: c.MustGet(\"db\").(*gorm.DB) failed",
			zap.String("错误码", strconv.FormatInt(int64(codes.GetDBError), 10)),
		)
		ResponseErrorWithCode(c, codes.GetDBError)
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		zap.L().Error("ProfileGet: c.Get(\"user_id\") failed",
			zap.String("错误码", strconv.FormatInt(int64(codes.GetUserIDError), 10)),
		)
		ResponseErrorWithCode(c, codes.GetUserIDError)
		return
	}

	userIDUint, ok := userID.(uint)
	if !ok {
		zap.L().Error("ProfileGet: userID.(uint) failed",
			zap.String("错误码", strconv.FormatInt(int64(codes.UserIDTypeError), 10)),
			zap.Any("用户ID", userID),
		)
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

// ProfilePut 用于处理用户修改个人信息的Put请求
// @Summary 修改个人信息
// @Description 修改个人信息
// @Tags 个人信息
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object body model.RegisterRequest true "用户信息"
// @Success 200 {object} controller.ResponseData "修改成功"
// @Failure 400 {object} controller.ResponseData "修改失败"
// @Router /profile/profile [put]
func ProfilePut(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		zap.L().Error("ProfilePut: c.MustGet(\"db\").(*gorm.DB) failed",
			zap.String("错误码", strconv.FormatInt(int64(codes.GetDBError), 10)),
		)
		ResponseErrorWithCode(c, codes.GetDBError)
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		zap.L().Error("ProfilePut: c.Get(\"user_id\") failed",
			zap.String("错误码", strconv.FormatInt(int64(codes.GetUserIDError), 10)),
		)
		ResponseErrorWithCode(c, codes.GetUserIDError)
		return
	}

	userIDUint, ok := userID.(uint)
	if !ok {
		zap.L().Error("ProfilePut: userID.(uint) failed",
			zap.String("错误码", strconv.FormatInt(int64(codes.UserIDTypeError), 10)),
			zap.Any("用户ID", userID),
		)
		ResponseErrorWithCode(c, codes.UserIDTypeError)
		return
	}

	var userProfileUpdateReq model.UserReq
	err := c.ShouldBind(&userProfileUpdateReq)
	// fmt.Print(userProfileUpdateReq)
	if err != nil {
		zap.L().Error("ProfilePut: c.ShouldBind(&userProfileUpdateReq ) failed",
			zap.Int("错误码", codes.BindDataError.Int()),
		)
		ResponseErrorWithCode(c, codes.BindDataError)
		return
	}

	apiError := logic.ModifyUserProfile(db, &userProfileUpdateReq, userIDUint)
	if apiError != nil {
		zap.L().Error("ProfilePut: logic.ModifyUserProfile failed",
			zap.Int("错误码", apiError.StatusCode.Int()),
			zap.Any("用户信息", userProfileUpdateReq),
		)
		ResponseError(c, *apiError)
		return
	}
	ResponseSuccess(c, nil)
	return
}

// ProfileAvatarPut 用于处理用户修改头像的Put请求
func ProfileAvatarPut(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		zap.L().Error("ProfilePut: c.MustGet(\"db\").(*gorm.DB) failed",
			zap.String("错误码", strconv.FormatInt(int64(codes.GetDBError), 10)),
		)
		ResponseErrorWithCode(c, codes.GetDBError)
		return
	}
	var avatar model.UserAvatarReq
	avatar.UserID = c.MustGet("user_id").(uint)
	err := c.ShouldBind(&avatar)
	if err != nil {
		zap.L().Error("ProfilePut: c.ShouldBind(&avatar) failed",
			zap.Int("错误码", codes.BindDataError.Int()),
		)
		ResponseErrorWithCode(c, codes.BindDataError)
		return
	}
	apiError := logic.ModifyUserAvatar(db, &avatar, c)
	if apiError != nil {
		zap.L().Error("ProfilePut: logic.ModifyUserAvatar failed",
			zap.Int("错误码", apiError.StatusCode.Int()),
			zap.Any("用户头像信息", avatar),
		)
		ResponseError(c, *apiError)
		return
	}
	ResponseSuccess(c, nil)
}
