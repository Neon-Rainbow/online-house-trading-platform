package router

import (
	"net/http"
	"online-house-trading-platform/internal/controller"
	"online-house-trading-platform/middleware"

	docs "online-house-trading-platform/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// setupUserAPI 建设了一个用户API,用于处理用户的个人信息,url为/user/:user_id
func setupUserAPI(r *gin.Engine, db *gorm.DB) {
	userGroup := r.Group("/user/:user_id").Use(
		middleware.JWTAuthMiddleware(),
		middleware.UserIDMatchMiddleware(),
		middleware.DBMiddleware(db))
	{
		userGroup.GET("/release", controller.ReleaseGet)
		userGroup.POST("/release", controller.ReleasePost)
		userGroup.GET("/favourites", controller.GetUserFavourites)
	}
	userProfileGroup := r.Group("/user/:user_id/profile").Use(
		middleware.JWTAuthMiddleware(),
		middleware.UserIDMatchMiddleware(),
		middleware.DBMiddleware(db))
	//TODO: 这里需要继续写
	{
		userProfileGroup.GET("/", controller.ProfileGet)
	}
}

func setupRootAPI(router *gin.Engine, db *gorm.DB) {
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
	setupRootAPI(router, db)

	//设置路由,地址为/auth
	setupAuthAPI(router, db)

	//设置路由,地址为/houses
	setupHouseAPI(router, db)

	//设置路由,地址为/user
	setupUserAPI(router, db)

	//404界面
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	docs.SwaggerInfo.BasePath = "/api/v1"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
