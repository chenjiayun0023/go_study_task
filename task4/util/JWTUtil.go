package util

import (
	"go_study/task4/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 自定义Claims
type CustomClaims struct {
	UserID   uint   `json:"userID"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 生成JWT令牌
func GenerateToken(userID uint, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		userID,
		username,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.Cfg.JWTExpire * time.Hour)), //24 * time.Hour  30 * time.Second
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})
	return token.SignedString([]byte(config.Cfg.JWTSecret))
}
