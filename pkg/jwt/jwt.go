package jwt

import (
	"online-house-trading-platform/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const tokenExpireDuration = time.Hour * 24 // token 过期时间

type MyClaims struct {
	Username string `json:"username"`
	UserID   uint   `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken 用于生成JWT
func GenerateToken(username string, userId uint) (string, error) {
	var jwtSecret = config.AppConfig.JWTSecret
	var mySecret = []byte(jwtSecret) // 自定义密钥
	c := MyClaims{
		Username: username,
		UserID:   userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpireDuration)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                          // 签发时间
			Issuer:    "409宿舍的精致的综合项目",                                         // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(mySecret)
}

// ParseToken 用于解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	var jwtSecret = config.AppConfig.JWTSecret
	var mySecret = []byte(jwtSecret) // 自定义密钥
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return mySecret, nil
		})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
