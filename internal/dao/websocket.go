package dao

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// InsertMsg 插入消息到 Redis，并设置相关字段和过期时间
func InsertMsg(rdb *redis.Client, id string, content string, status int, expireDuration int64) error {
	ctx := context.Background()

	// 创建一个复合键来存储消息
	key := "msg:" + id

	// 将消息的各个字段存储在 Redis 的哈希表中
	err := rdb.HSet(ctx, key, map[string]interface{}{
		"content": content,
		"status":  status,
		"created": time.Now().Unix(),
	}).Err()
	if err != nil {
		return err
	}

	// 设置过期时间
	err = rdb.Expire(ctx, key, time.Duration(expireDuration)*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}
