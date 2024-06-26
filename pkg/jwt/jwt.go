package jwt

import (
	"online-house-trading-platform/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	accessTokenExpireDuration  = time.Hour * 24 * 7 // 访问令牌过期时间
	refreshTokenExpireDuration = time.Hour * 24 * 7 // 刷新令牌过期时间
)

type MyClaims struct {
	Username  string `json:"username"`
	UserID    uint   `json:"user_id"`
	UserRole  string `json:"role"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}

// GenerateToken 用于生成JWT
// @title GenerateToken
// @description 生成JWT
// @param username string 用户名
// @param userId uint 用户ID
// @param role string 用户角色
// @return accessToken string 访问令牌
// @return refreshToken string 刷新令牌
// @return err error 错误信息
func GenerateToken(username string, userId uint, role string) (accessToken string, refreshToken string, err error) {
	var jwtSecret = config.AppConfig.JWTSecret
	var mySecret = []byte(jwtSecret) // 自定义密钥
	accessTokenClaims := MyClaims{
		Username:  username,
		UserID:    userId,
		UserRole:  role,
		TokenType: "access_token",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenExpireDuration)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                // 签发时间
			Issuer:    "409宿舍的精致的综合项目",                                               // 签发人
		},
	}
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims).SignedString(mySecret)
	if err != nil {
		return
	}

	refreshTokenClaims := MyClaims{
		Username:  username,
		UserID:    userId,
		UserRole:  role,
		TokenType: "refresh_token",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTokenExpireDuration)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                 // 签发时间
			Issuer:    "409宿舍的精致的综合项目",                                                // 签发人
		},
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims).SignedString(mySecret)
	if err != nil {
		return
	}
	return
}

// ParseToken 用于解析JWT
// @title ParseToken
// @description 解析JWT
// @param tokenString string JWT字符串
// @return *MyClaims JWT的Payload
// @return error 错误信息
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
