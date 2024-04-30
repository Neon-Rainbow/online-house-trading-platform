package logic

import (
	"crypto/md5"
	"encoding/hex"
	"online-house-trading-platform/config"
)

// EncryptPassword 用于对密码进行加密,使用了md5加密算法
func EncryptPassword(password string) string {
	// secret 用于存储密码加密的密钥
	var secret = config.AppConfig.PasswordSecret
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
