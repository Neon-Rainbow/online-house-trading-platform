package main

import (
	"fmt"
	"online-house-trading-platform/config"
	"online-house-trading-platform/logger" // 导入 logger 包
	"online-house-trading-platform/pkg/database"
	"online-house-trading-platform/pkg/redis"
	"online-house-trading-platform/router"

	"go.uber.org/zap"
)

// @title 在线房屋交易平台API文档
// @version 1.0
// @description 这是在线房屋交易平台的API文档, 用于提供房屋交易相关的接口, 包括用户注册、登录、房屋信息的增删改查等功能, 以及房屋图片的上传和获取等功能
// @termsOfService http://swagger.io/terms/

// @contact.name FHN
// @contact.email Haonan_Fang@Outlook.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	// 加载配置文件
	config.LoadConfig("config.json")

	// 初始化日志记录器
	logFilePath := config.AppConfig.LogFilePath
	logger := logger.InitLogger(logFilePath)
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	db := database.InitializeDB()
	if db == nil {
		zap.L().Error("数据库连接失败")
		return
	}

	_, err := redis.InitRedis()
	if err != nil {
		zap.L().Error("Redis连接失败", zap.Error(err))
	}

	route := router.SetupRouters(db)

	err = route.Run(fmt.Sprintf(":%d", config.AppConfig.Port))
	if err != nil {
		zap.L().Error("服务器连接失败", zap.Error(err))
		return
	}
}
