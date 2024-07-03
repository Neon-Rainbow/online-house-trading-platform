package controller

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/logic"
	"online-house-trading-platform/pkg/model"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	var loginReq model.LoginRequest

	err := c.ShouldBind(&loginReq)
	if err != nil {
		zap.L().Error("LoginPost: c.ShouldBind(&loginReq) failed",
			zap.Int("错误码", codes.LoginInvalidParam.Int()),
		)
		ResponseErrorWithCode(c, codes.LoginInvalidParam)
		return
	}

	resultChannel := make(chan *model.LoginResponse, 1)
	errorChannel := make(chan *model.Error, 1)

	go func() {
		defer close(resultChannel)
		defer close(errorChannel)
		loginResp, apiError := logic.LoginHandle(loginReq, c.Request.Header.Get("X-Real-IP"), c.Request.UserAgent())
		if apiError != nil {
			errorChannel <- apiError
			return
		}
		resultChannel <- loginResp
		return
	}()

	select {
	case loginResp := <-resultChannel:
		ResponseSuccess(c, loginResp)
	case apiError := <-errorChannel:
		zap.L().Error("LoginPost: logic.LoginHandle failed",
			zap.Int("错误码", apiError.StatusCode.Int()),
			zap.Any("loginReq", loginReq),
		)
		ResponseError(c, *apiError)
	case <-time.After(10 * time.Second):
		ResponseTimeout(c)
	}
	return

	//loginResp, apiError := logic.LoginHandle(loginReq, c.Request.Header.Get("X-Real-IP"), c.Request.UserAgent())
	//if apiError != nil {
	//	zap.L().Error("LoginPost: logic.LoginHandle failed",
	//		zap.Int("错误码", apiError.StatusCode.Int()),
	//		zap.Any("loginReq", loginReq),
	//	)
	//	ResponseError(c, *apiError)
	//	return
	//}
	//
	//ResponseSuccess(c, loginResp)
	//return
}

// AdminLogin 用于处理管理员的登录界面的POST请求
func AdminLogin(c *gin.Context) {
	var loginReq model.LoginRequest
	err := c.ShouldBind(&loginReq)
	if err != nil {
		zap.L().Error("LoginPost: c.ShouldBind(&loginReq) failed",
			zap.Int("错误码", codes.LoginInvalidParam.Int()),
		)
		ResponseErrorWithCode(c, codes.LoginInvalidParam)
		return
	}
	resultCh := make(chan *model.LoginResponse, 1)
	errorCh := make(chan *model.Error, 1)

	go func() {
		defer close(resultCh)
		defer close(errorCh)
		loginResp, apiError := logic.AdminLoginHandle(loginReq, c.Request.Header.Get("X-Real-IP"), c.Request.UserAgent())
		if apiError != nil {
			errorCh <- apiError
			return
		}
		resultCh <- loginResp
		return
	}()

	select {
	case loginResp := <-resultCh:
		ResponseSuccess(c, loginResp)
	case apiError := <-errorCh:
		zap.L().Error("LoginPost: logic.LoginHandle failed",
			zap.Int("错误码", apiError.StatusCode.Int()),
			zap.Any("loginReq", loginReq),
		)
		ResponseError(c, *apiError)
	case <-time.After(10 * time.Second):
		ResponseTimeout(c)
	}
	return

	//loginResp, apiError := logic.AdminLoginHandle(loginReq, c.Request.Header.Get("X-Real-IP"), c.Request.UserAgent())
	//if apiError != nil {
	//	zap.L().Error("LoginPost: logic.LoginHandle failed",
	//		zap.Int("错误码", apiError.StatusCode.Int()),
	//		zap.Any("loginReq", loginReq),
	//	)
	//	ResponseError(c, *apiError)
	//	return
	//}
	//
	//ResponseSuccess(c, loginResp)
	//return
}

// 获取客户端 IP 的函数
//func getClientIP(c *gin.Context) string {
//	// 优先从 X-Forwarded-For 头部获取
//	ip := c.Request.Header.Get("X-Forwarded-For")
//	if ip != "" {
//		// X-Forwarded-For 可能包含多个 IP 地址，用逗号分隔，取第一个
//		ips := strings.Split(ip, ",")
//		if len(ips) > 0 {
//			ip = strings.TrimSpace(ips[0])
//		}
//	}
//
//	// 如果 X-Forwarded-For 为空，则尝试从 X-Real-IP 头部获取
//	if ip == "" {
//		ip = c.Request.Header.Get("X-Real-IP")
//	}
//
//	// 如果 X-Real-IP 也为空，则从 RemoteAddr 获取
//	if ip == "" {
//		ip, _, _ = net.SplitHostPort(c.Request.RemoteAddr)
//	}
//
//	return ip
//}
