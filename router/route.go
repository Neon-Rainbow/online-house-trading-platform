package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"online-house-trading-platform/config"
	docs "online-house-trading-platform/docs"
	"online-house-trading-platform/internal/controller"
	"online-house-trading-platform/logger" // 导入 logger 包
	"online-house-trading-platform/middleware"
	"time"
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

	//router.Use(cors.Default())
	router.OPTIONS("/*path", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Status(http.StatusOK)
	})

	corsCFG := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(corsCFG))

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

	zap.L().Info("路由配置成功")
	return router
}
