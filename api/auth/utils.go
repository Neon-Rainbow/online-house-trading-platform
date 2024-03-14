package auth

import (
	"crypto/md5"
	"encoding/hex"
)

const secret = "409 Dormitory Project Practical Assignment"

// encryptPassword 用于对密码进行加密,使用了md5加密算法
func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
