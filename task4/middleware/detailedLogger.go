package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// bodyLogWriter 用于记录响应体
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// 详细接口日志中间件
func DetailedLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 跳过日志的路由（如健康检查）
		//if c.Request.URL.Path == "/health" {
		//	c.Next()
		//	return
		//}

		// 记录请求体
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 记录响应体
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		startTime := time.Now()
		requestID := c.GetString("X-Request-ID")

		// 记录请求开始
		log.Printf("[API-Start] %s | %s | %s %s | IP: %s | Body: %s",
			startTime.Format("2006-01-02 15:04:05"),
			requestID,
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			string(requestBody),
		)

		// 处理请求
		c.Next()

		// 记录请求结束
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// 记录响应
		var responseBody interface{}
		if blw.body.Len() > 0 {
			json.Unmarshal(blw.body.Bytes(), &responseBody)
		}

		log.Printf("[API-End] %s | %s | %d | %v | %s %s | Response: %v",
			endTime.Format("2006-01-02 15:04:05"),
			requestID,
			c.Writer.Status(),
			latency,
			c.Request.Method,
			c.Request.URL.Path,
			responseBody,
		)
	}
}
