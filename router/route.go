package router

import (
	"net/http"
	"online-house-trading-platform/internal/controller"
	"online-house-trading-platform/middleware"

	docs "online-house-trading-platform/docs"

	"online-house-trading-platform/logger" // 导入 logger 包

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func setupAuthAPI(r *gin.Engine, db *gorm.DB) {
	userGroup := r.Group("/auth").Use(middleware.DBMiddleware(db))
	{
		userGroup.POST("/login", controller.LoginPost)
		userGroup.POST("/register", controller.RegisterPost)
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

func setupUserAPI(r *gin.Engine, db *gorm.DB) {
	userGroup := r.Group("/user/:user_id").Use(
		middleware.JWTAuthMiddleware(),
		middleware.UserIDMatchMiddleware(),
		middleware.DBMiddleware(db))
	{
		userGroup.POST("/release", controller.ReleasePost)
		userGroup.PUT("/release", controller.ReleasePut)
		userGroup.DELETE("/release", controller.ReleaseDeleteWholeHouse)
		userGroup.GET("/favourites", controller.GetUserFavourites)
		userGroup.GET("/appointment", controller.HousesAppointmentGet)
	}
	userProfileGroup := r.Group("/user/:user_id/profile").Use(
		middleware.JWTAuthMiddleware(),
		middleware.UserIDMatchMiddleware(),
		middleware.DBMiddleware(db))
	{
		userProfileGroup.GET("/", controller.ProfileGet)
		userProfileGroup.PUT("/", controller.ProfilePut)
	}
}

// SetupRouters 设置web服务器路由
func SetupRouters(db *gorm.DB) *gin.Engine {
	router := gin.New()

	// 使用 zap 日志中间件
	router.Use(logger.GinLogger(zap.L()))
	router.Use(logger.GinRecovery(zap.L(), true))

	setupAuthAPI(router, db)
	setupHouseAPI(router, db)
	setupUserAPI(router, db)

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	docs.SwaggerInfo.BasePath = "/"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	zap.L().Info("路由配置成功")
	return router
}
