package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomePageGet 用来处理首页的Get请求
// @Summary 首页
// @Description 显示首页
// @Tags 首页
// @Accept json
// @Produce json
// @Success 200 {string} html "首页"
// @Router / [get]
func HomePageGet(c *gin.Context) {
	//c.HTML(http.StatusOK, "frontpage.html", nil)
	ResponseSuccess(c, nil)
}

// LearnMoreGet 用来处理LearnMore页面的Get请求
// @Summary LearnMore
// @Description 显示LearnMore页面
// @Tags LearnMore
// @Accept json
// @Produce json
// @Success 200 {string} html "LearnMore页面"
// @Failure 400 {object} controller.ResponseData "预约失败,具体原因查看json中的message字段和code字段"
// @Router /learn_more [get]
func LearnMoreGet(c *gin.Context) {
	c.HTML(http.StatusOK, "learn_more.html", nil)
}
