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

// RegisterGet 用于处理用户的注册界面的GET请求
// @Summary 注册界面
// @Description 显示用户注册界面
// @Tags 注册
// @Accept json
// @Produce json
// @Success 200 {string} html "注册界面"
// @Router /auth/register [get]
//func RegisterGet(c *gin.Context) {
//	//c.HTML(http.StatusOK, "register.html", nil)
//	ResponseSuccess(c, nil)
//	return
//}

// RegisterPost 用于处理用户的注册界面的POST请求
// @Summary 注册接口
// @Description 用户注册接口
// @Tags 注册
// @Accept json
// @Produce json
// @Param object query model.RegisterRequest false "查询参数"
// @Success 200 {object} controller.ResponseData "注册成功"
// @Failure 400 {object} controller.ResponseData "预约失败,具体原因查看json中的message字段和code字段"
// @Router /auth/register [post]
func RegisterPost(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		zap.L().Error("RegisterPost: c.MustGet(\"db\").(*gorm.DB) failed",
			zap.String("错误码", strconv.FormatInt(int64(codes.GetDBError), 10)),
		)
		ResponseErrorWithCode(c, codes.GetDBError)
		return
	}

	var registerReq model.RegisterRequest
	err := c.ShouldBind(&registerReq)
	registerReq.Role = "user"
	if err != nil {
		zap.L().Error("RegisterPost: c.ShouldBind(&registerReq) failed",
			zap.Int("错误码", codes.RegisterInvalidParam.Int()),
		)
		ResponseErrorWithCode(c, codes.RegisterInvalidParam)
		return
	}

	apiError := logic.RegisterHandle(db, registerReq, c)
	if apiError != nil {
		zap.L().Error("RegisterPost: logic.RegisterHandle failed",
			zap.Int("错误码", apiError.StatusCode.Int()),
			zap.Any("注册信息registerReq", registerReq),
		)
		ResponseError(c, *apiError)
		return
	}

	ResponseSuccess(c, nil)
	return
}

//func AdminRegisterPost(c *gin.Context) {
//	db, exist := c.MustGet("db").(*gorm.DB)
//	if !exist {
//		zap.L().Error("RegisterPost: c.MustGet(\"db\").(*gorm.DB) failed",
//			zap.String("错误码", strconv.FormatInt(int64(codes.GetDBError), 10)),
//		)
//		ResponseErrorWithCode(c, codes.GetDBError)
//		return
//	}
//
//	var registerReq model.AdminRegisterRequest
//	err := c.ShouldBind(&registerReq)
//	registerReq.Role = "admin"
//	if err != nil {
//		zap.L().Error("RegisterPost: c.ShouldBind(&registerReq) failed",
//			zap.Int("错误码", codes.RegisterInvalidParam.Int()),
//		)
//		ResponseErrorWithCode(c, codes.RegisterInvalidParam)
//		return
//	}
//	if registerReq.AdminSecret != config.AppConfig.AdminRegisterSecretKey {
//		zap.L().Error("RegisterPost: admin secret key error",
//			zap.Int("错误码", codes.RegisterInvalidParam.Int()),
//		)
//		ResponseErrorWithCode(c, codes.RegisterInvalidParam)
//		return
//	}
//
//	apiError := logic.RegisterHandle(db, registerReq, c)
//	if apiError != nil {
//		zap.L().Error("RegisterPost: logic.RegisterHandle failed",
//			zap.Int("错误码", apiError.StatusCode.Int()),
//			zap.Any("注册信息registerReq", registerReq),
//		)
//		ResponseError(c, *apiError)
//		return
//	}
//
//	ResponseSuccess(c, nil)
//	return
//}
