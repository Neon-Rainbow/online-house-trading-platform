package main

import (
	"log"
	"online-house-trading-platform/api"
	"online-house-trading-platform/config"
	"online-house-trading-platform/pkg/database"
)

// main 函数用于启动服务器,服务器端口为8080
func main() {
	config.LoadConfig("config.json") // 加载配置文件

	log.Println("Hello, World!")

	db := database.InitializeDB()
	if db == nil {
		log.Fatal("数据库初始化失败")
	}

	router := api.SetupRouter(db)

	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("服务器连接失败: %v", err)
	}
}
