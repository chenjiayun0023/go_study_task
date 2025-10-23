package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
)

// 自定义恢复中间件
func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 获取堆栈信息
				stack := debug.Stack()

				// 记录错误日志
				log.Printf("[Recovery] %s Panic recovered:\nError: %v\nStack: %s\n",
					time.Now().Format("2006-01-02 15:04:05"),
					err,
					string(stack),
				)

				// 返回统一错误响应
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":      "服务器内部错误",
					"code":       "INTERNAL_SERVER_ERROR",
					"request_id": c.GetString("X-Request-ID"),
					"timestamp":  time.Now().Unix(),
				})

				c.Abort()
			}
		}()
		c.Next()
	}
}
