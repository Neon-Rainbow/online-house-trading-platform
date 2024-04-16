package main

import (
	"fmt"
	"log"
	"online-house-trading-platform/config"
	"online-house-trading-platform/logger"
	"online-house-trading-platform/pkg/database"
	"online-house-trading-platform/router"
)

// main 函数用于启动服务器,服务器端口为8080
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
