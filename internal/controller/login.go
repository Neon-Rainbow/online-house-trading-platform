package controller

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// LoginGet 用于处理用户的登录界面的GET请求
// @Summary 登录界面
// @Description 显示用户登录界面
// @Tags 登录
// @Accept json
// @Produce json
// @Success 200 {string} html "登录界面"
// @Router /auth/login [get]
//func LoginGet(c *gin.Context) {
//	//c.HTML(http.StatusOK, "login.html", nil)
//	ResponseSuccess(c, nil)
//	return
//}

// LoginPost 用于处理用户的登录界面的POST请求
// @Summary 登录接口
// @Description 用户登录接口
// @Tags 登录
// @Accept json
// @Produce json
// @Param object query model.LoginRequest false "查询参数"
// @Success 200 {object} controller.ResponseData "登录成功"
// @Failure 400 {object} controller.ResponseData "预约失败,具体原因查看json中的message字段和code字段"
// @Router /auth/login [post]
func LoginPost(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		zap.L().Error("LoginPost: c.MustGet(\"db\").(*gorm.DB) failed")
		ResponseErrorWithCode(c, codes.GetDBError)
		return
	}

	var loginReq model.LoginRequest

	err := c.ShouldBind(&loginReq)
	if err != nil {
		zap.L().Error("LoginPost: c.ShouldBind(&loginReq) failed",
			zap.Int("错误码", codes.LoginInvalidParam.Int()),
		)
		ResponseErrorWithCode(c, codes.LoginInvalidParam)
		return
	}

	loginResp, apiError := logic.LoginHandle(db, loginReq, c)
	if apiError != nil {
		zap.L().Error("LoginPost: logic.LoginHandle failed",
			zap.Int("错误码", apiError.StatusCode.Int()),
			zap.Any("loginReq", loginReq),
		)
		ResponseError(c, *apiError)
		return
	}

	ResponseSuccess(c, loginResp)
	return
}

// AdminLogin 用于处理管理员的登录界面的POST请求
func AdminLogin(c *gin.Context) {
	db, exist := c.MustGet("db").(*gorm.DB)
	if !exist {
		zap.L().Error("LoginPost: c.MustGet(\"db\").(*gorm.DB) failed")
		ResponseErrorWithCode(c, codes.GetDBError)
		return
	}

	var loginReq model.LoginRequest
	err := c.ShouldBind(&loginReq)
	if err != nil {
		zap.L().Error("LoginPost: c.ShouldBind(&loginReq) failed",
			zap.Int("错误码", codes.LoginInvalidParam.Int()),
		)
		ResponseErrorWithCode(c, codes.LoginInvalidParam)
		return
	}
	loginResp, apiError := logic.AdminLoginHandle(db, loginReq, c)
	if apiError != nil {
		zap.L().Error("LoginPost: logic.LoginHandle failed",
			zap.Int("错误码", apiError.StatusCode.Int()),
			zap.Any("loginReq", loginReq),
		)
		ResponseError(c, *apiError)
		return
	}
	ResponseSuccess(c, loginResp)
	return
}
