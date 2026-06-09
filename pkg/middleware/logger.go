package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	"sanguo-server/pkg/logger"
)

// AccessLog 记录 HTTP 请求日志
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		logger.WithFields(map[string]interface{}{
			"status":  c.Writer.Status(),
			"method":  c.Request.Method,
			"path":    path,
			"latency": time.Since(start).String(),
			"ip":      c.ClientIP(),
		}).Info("access")
	}
}
