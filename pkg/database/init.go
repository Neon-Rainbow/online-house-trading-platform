package database

import (
	"fmt"
	"log"
	"online-house-trading-platform/config"
	"online-house-trading-platform/pkg/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database 用于存储数据库连接
var Database *gorm.DB

// InitializeDB 初始化数据库并进行迁移
// @title InitializeDB
// @description 初始化数据库并进行迁移
// @return db *gorm.DB 数据库连接
// @return err error 错误信息
func InitializeDB() (db *gorm.DB, err error) {

	// 数据库连接字符串，替换为你的实际数据库信息
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.AppConfig.Database.User,
		config.AppConfig.Database.Password,
		config.AppConfig.Database.Host,
		config.AppConfig.Database.Port,
		config.AppConfig.Database.DBName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("服务器连接失败: %v", err)
		return nil, err
	}

	// 使用AutoMigrate迁移模型
	err = db.AutoMigrate(
		&model.User{},
		&model.House{},
		&model.HouseImage{},
		&model.Favourite{},
		&model.Reserve{},
		&model.UserAvatar{},
		&model.LoginRecord{},
		&model.ViewingRecords{})
	if err != nil {
		return nil, err
	}

	log.Printf("数据库连接成功")

	// 将数据库连接赋值给全局变量Database
	Database = db

	return db, nil
}
