package controller

import (
	"online-house-trading-platform/codes"
	"online-house-trading-platform/pkg/jwt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var tokenBlacklist = struct {
	sync.RWMutex
	tokens map[string]time.Time
}{tokens: make(map[string]time.Time)}

func addToBlacklist(token string, expiration time.Time) {
	tokenBlacklist.Lock()
	tokenBlacklist.tokens[token] = expiration
	tokenBlacklist.Unlock()
}

func isTokenBlacklisted(token string) bool {
	tokenBlacklist.RLock()
	expiration, exists := tokenBlacklist.tokens[token]
	tokenBlacklist.RUnlock()

	// 检查令牌是否存在以及是否已经过期
	if exists && time.Now().Before(expiration) {
		return true
	}
	return false
}

// RefreshToken 用于刷新令牌
// @Summary 刷新令牌
// @Description 刷新令牌
// @Tags 用户
// @Accept json
// @Produce json
// @Param refresh_token query string true "刷新令牌"
// @Success 200 {object} Response
// @Router /refresh_token [get]
func RefreshToken(c *gin.Context) {
	refreshToken := c.Query("refresh_token") // 假设通过查询参数传递

	if isTokenBlacklisted(refreshToken) {
		ResponseErrorWithCode(c, codes.TokenBlacklistedError)
		return
	}

	claims, err := jwt.ParseToken(refreshToken)
	if err != nil {
		ResponseErrorWithCode(c, codes.InvalidTokenError)
		return
	}

	// 确保令牌未过期且为刷新令牌
	if time.Now().After(claims.ExpiresAt.Time) {
		ResponseErrorWithCode(c, codes.TokenExpiredError)
		return
	}

	newAccessToken, newRefreshToken, err := jwt.GenerateToken(claims.Username, claims.UserID, claims.UserRole)

	// 将旧的刷新令牌加入黑名单
	addToBlacklist(refreshToken, time.Now().Add(48*time.Hour)) // 假定黑名单上的令牌在48小时后清除

	if err != nil {
		ResponseErrorWithCode(c, codes.GenerateJWTTokenError)
		return
	}
	data := gin.H{
		"username":     claims.Username,
		"userID":       claims.UserID,
		"userRole":     claims.UserRole,
		"accessToken":  newAccessToken,
		"refreshToken": newRefreshToken,
	}
	ResponseSuccess(c, data)
}
