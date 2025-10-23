package middleware

import (
	"fmt"
	"go_study/task4/config"
	"go_study/task4/util"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// 鉴权中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

		token, err := jwt.ParseWithClaims(tokenString, &util.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.Cfg.JWTSecret), nil
		})
		if err != nil {
			log.Printf("token解析失败: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		if claims, ok := token.Claims.(*util.CustomClaims); ok && token.Valid {
			// 将用户信息存入上下文，供后续处理使用
			c.Set("userID", claims.UserID)
			c.Set("username", claims.Username)
			c.Set("exp", claims.ExpiresAt)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		}
	}
}
