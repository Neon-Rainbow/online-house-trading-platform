package controller

import (
	"net/http"
	"online-house-trading-platform/codes"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
)

// ResponseData 用于封装API的返回数据
type ResponseData struct {
	Code    codes.ResCode `json:"code"`    // 自定义的返回码
	Message interface{}   `json:"message"` // 返回的信息
	Data    interface{}   `json:"data"`    // 返回的数据
}

// ResponseSuccess 用于返回成功信息
func ResponseSuccess(c *gin.Context, data interface{}) {
	responseData := &ResponseData{
		Code:    codes.CodeSuccess,
		Message: codes.CodeSuccess.Message(),
		Data:    data,
	}
	c.JSON(http.StatusOK, responseData)
}

// ResponseError 用于返回错误信息
func ResponseError(c *gin.Context, error model.Error) {
	var msg string
	if error.Message == "" {
		msg = error.Message
	} else {
		msg = error.StatusCode.Message()
	}
	responseData := &ResponseData{
		Code:    error.StatusCode,
		Message: msg,
		Data:    nil,
	}
	c.JSON(http.StatusOK, responseData)
}

func ResponseErrorWithCode(c *gin.Context, code codes.ResCode) {
	responseData := &ResponseData{
		Code:    code,
		Message: code.Message(),
		Data:    nil,
	}
	c.JSON(http.StatusBadRequest, responseData)
}
