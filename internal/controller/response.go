package controller

import (
	"net/http"
	"online-house-trading-platform/pkg/model"

	"github.com/gin-gonic/gin"
)

// ResponseData 用于封装API的返回数据
type ResponseData struct {
	Code    ResCode     `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

// ResponseSuccess 用于返回成功信息
func ResponseSuccess(c *gin.Context, data interface{}) {
	responseData := &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.Message(),
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

func ResponseErrorWithCode(c *gin.Context, code ResCode) {
	responseData := &ResponseData{
		Code:    code,
		Message: code.Message(),
		Data:    nil,
	}
	c.JSON(http.StatusOK, responseData)
}
