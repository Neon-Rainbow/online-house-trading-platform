package router

import (
	"net/http"
	"online-house-trading-platform/config"
	docs "online-house-trading-platform/docs"
	"online-house-trading-platform/internal/controller"
	"online-house-trading-platform/logger" // 导入 logger 包
	"online-house-trading-platform/middleware"
	"time"

	"github.com/gin-contrib/cors"
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
		housesGroup.GET("/:house_id", controller.GetHouseInfoByID)
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
		userGroup.GET("/release", controller.ReleaseGet)
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
		userProfileGroup.PUT("/avatar", controller.ProfileAvatarPut)
	}
}

// SetupRouters 设置web服务器路由
func SetupRouters(db *gorm.DB) *gin.Engine {
	router := gin.New()
	gin.SetMode(config.AppConfig.GinMode)

	// 使用 zap 日志中间件
	router.Use(logger.GinLogger(zap.L()))
	router.Use(logger.GinRecovery(zap.L(), true))

	// 配置CORS跨域请求
	corsConfig := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowAllOrigins:  true,
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(corsConfig))

	setupAuthAPI(router, db)
	setupHouseAPI(router, db)
	setupUserAPI(router, db)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"url":     c.Request.RequestURI,
			"message": "无法访问",
		})
	})

	docs.SwaggerInfo.BasePath = "/"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//zap.L().Info("路由配置成功")
	return router
}
