package config

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

type Redis struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type Config struct {
	Database               Database `json:"database"`
	Redis                  Redis    `json:"redis"`
	JWTSecret              string   `json:"jwtSecret"`
	PasswordSecret         string   `json:"passwordSecret"`
	LogFilePath            string   `json:"logFilePath"`
	Port                   int      `json:"port"`
	GinMode                string   `json:"ginMode"`
	ZapLogLever            string   `json:"zapLogLever"`
	AdminRegisterSecretKey string   `json:"admin_register_secret_key"`
}

// AppConfig 用于存储配置文件的内容
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

	overrideWithEnvVariables(AppConfig)
}

// overrideWithEnvVariables 用于用环境变量覆盖配置文件中的值, 以便在容器中使用
func overrideWithEnvVariables(config *Config) {
	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		config.Database.Host = dbHost
	}
	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
		if port, err := strconv.Atoi(dbPort); err == nil {
			config.Database.Port = port
		}
	}
	if dbUser := os.Getenv("DB_USER"); dbUser != "" {
		config.Database.User = dbUser
	}
	if dbPassword := os.Getenv("DB_PASSWORD"); dbPassword != "" {
		config.Database.Password = dbPassword
	}
	if dbName := os.Getenv("DB_NAME"); dbName != "" {
		config.Database.DBName = dbName
	}
	if redisHost := os.Getenv("REDIS_HOST"); redisHost != "" {
		config.Redis.Host = redisHost
	}
	if redisPort := os.Getenv("REDIS_PORT"); redisPort != "" {
		if port, err := strconv.Atoi(redisPort); err == nil {
			config.Redis.Port = port
		}
	}
	if redisPassword := os.Getenv("REDIS_PASSWORD"); redisPassword != "" {
		config.Redis.Password = redisPassword
	}
}
