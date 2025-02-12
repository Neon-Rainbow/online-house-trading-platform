package controller

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"
	"online-house-trading-platform/pkg/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetUserProfileByUserID 用于处理用户获取个人信息的Get请求
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
func GetUserProfileByUserID(c *gin.Context) {
	userID := c.Param("user_id")

	userIDUint, _ := strconv.ParseUint(userID, 10, 64)

	userProfile, apiError := logic.GetUserProfile(uint(userIDUint))
	if apiError != nil {
		ResponseError(c, *apiError)
		return
	}
	ResponseSuccess(c, userProfile)
}

// UpdateUserProfileByUserID 用于处理用户修改个人信息的Put请求
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
func UpdateUserProfileByUserID(c *gin.Context) {
	userID := c.Param("user_id")
	userIDUint64, _ := strconv.ParseUint(userID, 10, 64)
	userIDUint := uint(userIDUint64)
	//if !ok {
	//	zap.L().Error("UpdateUserProfileByUserID: userID.(uint) failed",
	//		zap.String("错误码", strconv.FormatInt(int64(codes.UserIDTypeError), 10)),
	//		zap.Any("用户ID", userID),
	//	)
	//	ResponseErrorWithCode(c, codes.UserIDTypeError)
	//	return
	//}

	var userProfileUpdateReq model.UserReq
	err := c.ShouldBind(&userProfileUpdateReq)
	if userProfileUpdateReq.Password != "" {
		userProfileUpdateReq.Password = logic.EncryptPassword(userProfileUpdateReq.Password) //对需要修改的明文密码进行加密
	} else {
		temp, _ := logic.GetUserProfile(userIDUint)
		userProfileUpdateReq.Password = temp.Password
	}

	// fmt.Print(userProfileUpdateReq)
	if err != nil {
		zap.L().Error("UpdateUserProfileByUserID: c.ShouldBind(&userProfileUpdateReq ) failed",
			zap.Int("错误码", codes.BindDataError.Int()),
		)
		ResponseErrorWithCode(c, codes.BindDataError)
		return
	}

	apiError := logic.ModifyUserProfile(&userProfileUpdateReq, userIDUint)
	if apiError != nil {
		zap.L().Error("UpdateUserProfileByUserID: logic.ModifyUserProfile failed",
			zap.Int("错误码", apiError.StatusCode.Int()),
			zap.Any("用户信息", userProfileUpdateReq),
		)
		ResponseError(c, *apiError)
		return
	}
	ResponseSuccess(c, nil)
	return
}

// UpdateUserAvatarByUserID 用于处理用户修改头像的Put请求
func UpdateUserAvatarByUserID(c *gin.Context) {
	var avatar model.UserAvatarReq
	avatar.UserID = c.MustGet("user_id").(uint)
	err := c.ShouldBind(&avatar)
	if err != nil {
		zap.L().Error("UpdateUserProfileByUserID: c.ShouldBind(&avatar) failed",
			zap.Int("错误码", codes.BindDataError.Int()),
		)
		ResponseErrorWithCode(c, codes.BindDataError)
		return
	}
	apiError := logic.ModifyUserAvatar(&avatar)
	if apiError != nil {
		zap.L().Error("UpdateUserProfileByUserID: logic.ModifyUserAvatar failed",
			zap.Int("错误码", apiError.StatusCode.Int()),
			zap.Any("用户头像信息", avatar),
		)
		ResponseError(c, *apiError)
		return
	}
	ResponseSuccess(c, nil)
}
