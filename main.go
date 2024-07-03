package main

import (
	"fmt"
	"go.uber.org/zap"
	"online-house-trading-platform/config"
	"online-house-trading-platform/logger"
	"online-house-trading-platform/pkg/database"
	"online-house-trading-platform/pkg/redis"
	"online-house-trading-platform/router"
	"sync"
)

// @title 在线房屋交易平台API文档
// @version 1.0
// @description 这是在线房屋交易平台的API文档, 用于提供房屋交易相关的接口, 包括用户注册、登录、房屋信息的增删改查等功能, 以及房屋图片的上传和获取等功能
// @termsOfService http://swagger.io/terms/

// @contact.name FHN
// @contact.email Haonan_Fang@Outlook.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 0.0.0.0: config.AppConfig.Port
// @BasePath /
func main() {
	err := config.LoadConfig("config.json")
	if err != nil {
		zap.L().Error("加载配置文件失败", zap.Error(err))
		return
	}

	logFilePath := config.AppConfig.LogFilePath
	logger := logger.InitLogger(logFilePath)
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	errorChannel := make(chan error, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		_, err := database.InitializeDB()
		if err != nil {
			errorChannel <- err
		}
	}()

	go func() {
		defer wg.Done()
		_, err := redis.InitRedis()
		if err != nil {
			errorChannel <- err
		}
	}()

	go func() {
		wg.Wait()
		close(errorChannel)
	}()

	for err := range errorChannel {
		if err != nil {
			zap.L().Error("初始化数据库或Redis失败", zap.Error(err))
			return
		}
	}

	route := router.SetupRouters()

	err = route.Run(fmt.Sprintf("%s:%d", config.AppConfig.Address, config.AppConfig.Port))
	if err != nil {
		zap.L().Error("服务器连接失败", zap.Error(err))
		return
	}
}
