package middleware

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// 自定义日志中间件
func CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 生成请求ID
		requestID := fmt.Sprintf("%d", startTime.UnixNano())
		c.Set("X-Request-ID", requestID)

		// 记录请求日志
		log.Printf("[Request] %s | %s | %s | %s | %s",
			startTime.Format("2006-01-02 15:04:05"),
			requestID,
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
		)

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// 状态码
		statusCode := c.Writer.Status()

		// 记录响应日志
		log.Printf("[Response] %s | %s | %d | %v | %s | %s",
			endTime.Format("2006-01-02 15:04:05"),
			requestID,
			statusCode,
			latency,
			c.Request.Method,
			c.Request.URL.Path,
		)

		// 如果发生错误，记录错误日志
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				log.Printf("[Error] %s | %s | %s | %v",
					time.Now().Format("2006-01-02 15:04:05"),
					requestID,
					c.Request.URL.Path,
					e.Error(),
				)
			}
		}
	}
}
