package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
	} `json:"database"`
	JWTSecret      string `json:"jwtSecret"`
	PasswordSecret string `json:"passwordSecret"`
	LogFilePath    string `json:"logFilePath"`
	Port           int    `json:"port"`
}

var AppConfig *Config

// LoadConfig 用于加载配置文件
func LoadConfig(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("无法打开配置文件: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("关闭文件失败: %v", err)
		}
	}(file)

	AppConfig = &Config{}
	if err := json.NewDecoder(file).Decode(AppConfig); err != nil {
		log.Fatalf("无法解析配置文件: %v", err)
	}
}
