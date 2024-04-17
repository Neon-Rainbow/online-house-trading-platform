package main

import (
	"fmt"
	"log"
	"online-house-trading-platform/config"
	"online-house-trading-platform/logger"
	"online-house-trading-platform/pkg/database"
	"online-house-trading-platform/router"
)

// @title 在线房屋交易平台API文档
// @version 1.0
// @description 这是在线房屋交易平台的API文档, 用于提供房屋交易相关的接口, 包括用户注册、登录、房屋信息的增删改查等功能, 以及房屋图片的上传和获取等功能
// @termsOfService http://swagger.io/terms/

// @contact.name FHN
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
func main() {
	config.LoadConfig("config.json") // 加载配置文件

	// 打开日志文件
	logFile := logger.InitLogger(config.AppConfig.LogFilePath)
	defer func() {
		err := logFile.Close()
		if err != nil {
			log.Fatalf("关闭日志文件失败: %v", err)
		}
	}()

	log.Println("项目启动")

	db := database.InitializeDB()
	if db == nil {
		log.Fatal("数据库初始化失败")
	}

	route := router.SetupRouters(db)

	err := route.Run(fmt.Sprintf(":%d", config.AppConfig.Port))

	if err != nil {
		log.Fatalf("服务器连接失败: %v", err)
	}
}
