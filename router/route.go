package router

import (
	"net/http"
	"online-house-trading-platform/internal/controller"
	"online-house-trading-platform/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// setupAuthAPI 建设了一个用户API,用于处理用户的登录和注册,url为/auth
// 该API包含了以下功能:
// 1. 用户登录,其中GET请求用于获取登录界面,POST请求用于处理登录请求
// 2. 用户注册,其中GET请求用于获取注册界面,POST请求用于处理注册请求
// 3. 用户登出,其中POST请求用于处理登出请求
// 该API使用了数据库中间件,用于获取数据库连接
func setupAuthAPI(r *gin.Engine, db *gorm.DB) {
	userGroup := r.Group("/auth").Use(middleware.DBMiddleware(db))
	{
		userGroup.GET("/login", controller.LoginGet)
		userGroup.POST("/login", controller.LoginPost)

		userGroup.GET("/register", controller.RegisterGet)
		userGroup.POST("/register", controller.RegisterPost)

		// userGroup.POST("/logout", controller.LogoutPost)
	}
}

func setupHouseAPI(r *gin.Engine, db *gorm.DB) {
	housesGroup := r.Group("/houses").Use(middleware.DBMiddleware(db))
	{
		housesGroup.GET("/", controller.GetAllHouses)
		housesGroup.GET("/:id", controller.GetHouseInfoByID)
		housesGroup.POST("/appointment", middleware.JWTAuthMiddleware(), controller.HousesAppointmentPost)
		housesGroup.POST("/collect", middleware.JWTAuthMiddleware(), controller.CollectPost)
	}
}

func SetupRootAPI(router *gin.Engine, db *gorm.DB) {
	router.GET("/", controller.HomePageGet)
	router.GET("/learn_more", controller.LearnMoreGet)
}

// SetupRouters 设置web服务器路由
func SetupRouters(db *gorm.DB) *gin.Engine {

	router := gin.Default()

	//加载静态文件
	router.Static("/static", "./web/static")

	//加载模板文件
	router.LoadHTMLGlob("./web/templates/**/*")

	//设置路由,地址为/
	SetupRootAPI(router, db)

	//设置路由,地址为/auth
	setupAuthAPI(router, db)

	//设置路由,地址为/houses
	setupHouseAPI(router, db)

	//404界面
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	return router
}
