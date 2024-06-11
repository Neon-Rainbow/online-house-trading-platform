package controller

import (
	"fmt"
	"online-house-trading-platform/codes"

	"github.com/gin-gonic/gin"
)

type fileUrl struct {
	Url string `json:"url" form:"url"`
}

// GetFileByURL 处理根据URL获取文件的请求
// @Summary 根据URL获取文件
// @Description 通过提供的URL获取文件
// @Tags 文件
// @Accept json
// @Produce application/octet-stream
// @Param url query string true "文件URL"
// @Success 200 {string} string "请求成功"
// @Failure 400 {object} object "请求参数错误"
// @Router /getFile [get]
func GetFileByURL(c *gin.Context) {
	var url fileUrl
	err := c.ShouldBind(&url)
	if err != nil {
		ResponseErrorWithCode(c, codes.LoginInvalidParam)
		return
	}
	filePath := url.Url
	if filePath == "" {
		ResponseErrorWithCode(c, codes.LoginInvalidParam)
	}

	fmt.Println("filepath:" + filePath)
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filePath))
	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	c.File(filePath)
	ResponseSuccess(c, filePath)
	return
}
