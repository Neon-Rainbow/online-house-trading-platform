package database

import (
	"encoding/json"
	"fmt"
	"log"
	"online-house-trading-platform/config"
	"online-house-trading-platform/pkg/model"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
	} `json:"database"`
}

// LoadConfig 用于加载配置文件中的数据库配置
func loadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("关闭文件失败: %v", err)
		}
	}(file)

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// InitializeDB 初始化数据库并进行迁移
func InitializeDB() *gorm.DB {

	// 数据库连接字符串，替换为你的实际数据库信息
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.AppConfig.Database.User,
		config.AppConfig.Database.Password,
		config.AppConfig.Database.Host,
		config.AppConfig.Database.Port,
		config.AppConfig.Database.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("服务器连接失败: %v", err)
	}

	// 使用AutoMigrate迁移模型
	err = db.AutoMigrate(&model.User{}, &model.House{}, &model.HouseImage{}, &model.Favourite{}, &model.Reserve{})
	if err != nil {
		return nil
	}
	
	log.Printf("数据库连接成功")

	return db
}
