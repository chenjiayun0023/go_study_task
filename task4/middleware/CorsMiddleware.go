package middleware

import "github.com/gin-gonic/gin"

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 在请求处理前设置CORS头
		c.Header("Access-Control-Allow-Origin", "http://localhost:8087")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, X-Requested-With, X-Request-ID")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, X-Request-ID")
		c.Header("Access-Control-Max-Age", "86400") // 24小时
		c.Header("Access-Control-Allow-Credentials", "true")

		// 2. 处理OPTIONS预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // 直接返回，不继续后续中间件
			return
		}

		// 3. 传递给下一个中间件
		c.Next()

		// 4. 在响应返回前可以再次处理CORS头（如果需要）
		// 这里通常不需要额外处理，因为已经在前面设置了
	}
}
