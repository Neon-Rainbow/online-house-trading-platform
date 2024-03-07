package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"online-house-trading-platform/api/auth"
	"online-house-trading-platform/api/houses"
	"online-house-trading-platform/api/user"
)

// SetupRouter 设置web服务器路由
func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	//加载静态文件
	router.Static("/static", "static")

	//加载模板文件
	router.LoadHTMLGlob("static/**/*")

	//设置路由,地址为/auth
	auth.SetUpAuthAPI(router, db)

	//设置路由,地址为/user
	user.SetUpUserAPI(router, db)

	//设置路由,地址为/houses
	houses.SetUpHousesAPI(router, db)

	return router
}