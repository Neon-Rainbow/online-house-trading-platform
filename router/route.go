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
)

func setupAuthAPI(r *gin.Engine) {
	userGroup := r.Group("/auth")
	{
		userGroup.POST("/login", controller.LoginPost)
		userGroup.POST("/admin_login", controller.AdminLogin)
		userGroup.POST("/register", controller.RegisterPost)
	}
}

func setupHouseAPI(r *gin.Engine) {
	housesGroup := r.Group("/houses")
	{
		housesGroup.GET("/", controller.GetAllHouses)
		housesGroup.GET("/:house_id", middleware.JWTAuthMiddleware(), controller.GetHouseInformationByHouseID)
		housesGroup.POST("/appointment", middleware.JWTAuthMiddleware(), controller.HousesAppointmentPost)
		housesGroup.POST("/collect", middleware.JWTAuthMiddleware(), controller.CollectPost)
	}
}

func setupUserAPI(r *gin.Engine) {
	userGroup := r.Group("/user/:user_id").Use(
		middleware.JWTAuthMiddleware(),
		middleware.UserIDMatchMiddleware())
	{
		userGroup.GET("/release", controller.ReleaseGet)
		userGroup.POST("/release", controller.ReleasePost)
		userGroup.POST("/update_house", controller.ReleasePut)
		userGroup.POST("/delete_house", controller.DeleteHouseInformationByHouseID)
		userGroup.GET("/favourites", controller.GetUserFavouritesByUserID)
		userGroup.POST("/delete_favourite", controller.DeleteFavouriteByUserID)
		userGroup.GET("/appointment", controller.GetUserAppointmentsByUserID)
		userGroup.POST("/delete_appointment", controller.DeleteAppointment)
		userGroup.POST("/delete_account", controller.DeleteUserAccountByUserID)
		userGroup.GET("/get_login_record", controller.GetUserLoginRecordByUserID)
		userGroup.GET("/get_login_record_excel", controller.GetLoginRecordToExcel)
		userGroup.GET("/get_viewing_record", controller.GetUserViewingRecordsByUserID)
	}
	userProfileGroup := r.Group("/user/:user_id/profile").Use(
		middleware.JWTAuthMiddleware(),
		middleware.UserIDMatchMiddleware())
	{
		userProfileGroup.GET("/", controller.GetUserProfileByUserID)
		userProfileGroup.PUT("/", controller.UpdateUserProfileByUserID)
		userProfileGroup.PUT("/avatar", controller.UpdateUserAvatarByUserID)
	}
}

// setupAdminAPI 设置管理员API
func setupAdminAPI(r *gin.Engine) {
	adminGroup := r.Group("/admin").Use(
		middleware.JWTAuthMiddleware(),
		middleware.AdminMiddleWare(),
	)
	{
		//adminGroup.POST("/register", controller.AdminRegisterPost)

		adminGroup.GET("/users", controller.GetAllUsersInformation)
		adminGroup.GET("/users/:user_id", controller.GetUserProfileByUserID)
		adminGroup.POST("/users/:user_id/update_user_profile", controller.UpdateUserProfileByUserID)
		adminGroup.POST("/users/:user_id/update_user_avatar", controller.UpdateUserAvatarByUserID)
		adminGroup.POST("/users/:user_id/delete_user_account", controller.DeleteUserAccountByUserID)

		adminGroup.GET("/houses", controller.GetAllHousesInformation)
		adminGroup.GET("/houses/:house_id", controller.GetHouseInformationByHouseID)
		adminGroup.POST("/houses/:house_id/delete_house", controller.DeleteHouseInformationByHouseID)

		adminGroup.GET("/appointments", controller.GetAllAppointments)
		adminGroup.GET("/appointments/:user_id", controller.GetUserAppointmentsByUserID)

		adminGroup.GET("/favourites", controller.GetAllFavourites)
		adminGroup.GET("/favourites/:user_id", controller.GetUserFavouritesByUserID)

		adminGroup.GET("/login_records", controller.GetAllLoginRecords)
		adminGroup.GET("/login_records/:user_id", controller.GetUserLoginRecordByUserID)

		adminGroup.GET("/get_log_file", controller.GetLogFile)
		adminGroup.POST("/delete_log_file", controller.DeleteLogFile)
	}
}

func setupOtherRouter(r *gin.Engine) {
	r.GET("/getFile", controller.GetFileByURL)
	r.GET("/refresh_token", controller.RefreshToken)
}

// SetupRouters 设置web服务器路由
func SetupRouters() *gin.Engine {
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

	setupAuthAPI(router)
	setupHouseAPI(router)
	setupUserAPI(router)
	setupAdminAPI(router)
	setupOtherRouter(router)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"url":     c.Request.RequestURI,
			"message": "无法访问",
			"method":  c.Request.Method,
		})
	})

	docs.SwaggerInfo.BasePath = "/"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//manager := &logic.ClientManager{
	//	Clients:    make(map[string]*logic.Client),
	//	Broadcast:  make(chan *logic.Broadcast),
	//	Register:   make(chan *logic.Client),
	//	Reply:      make(chan *logic.Client),
	//	Unregister: make(chan *logic.Client),
	//}
	//
	//go manager.Start()

	//router.GET("/ws", controller.WebsocketHandler)

	return router
}
