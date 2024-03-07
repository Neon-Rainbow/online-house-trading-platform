package main

import (
	"log"
	"online-house-trading-platform/api"
	"online-house-trading-platform/pkg/database"
)

func main() {
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
