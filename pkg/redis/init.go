package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"online-house-trading-platform/config"
	"strconv"
)

var Rdb *redis.Client

// InitRedis 用于初始化Redis配置
func InitRedis() (*redis.Client, error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.AppConfig.Redis.Host, strconv.Itoa(config.AppConfig.Redis.Port)),
		Password: config.AppConfig.Redis.Password,
		DB:       config.AppConfig.Redis.DB,
	})

	_, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return Rdb, nil
}
