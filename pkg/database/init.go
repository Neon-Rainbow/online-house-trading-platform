package database

import (
	"log"
	"online-house-trading-platform/pkg/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitializeDB 初始化数据库并进行迁移
func InitializeDB() *gorm.DB {
	// 数据库连接字符串，替换为你的实际数据库信息
	dsn := "ubuntu:FHn20010930@tcp(124.223.10.155:3306)/online_house_trading?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("服务器连接失败: %v", err)
	}

	// 使用AutoMigrate迁移模型
	err = db.AutoMigrate(&model.User{}, &model.House{}, &model.Favourite{})
	if err != nil {
		return nil
	}

	return db
}
