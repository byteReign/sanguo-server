package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"

	"sanguo-server/pkg/errors"
	"sanguo-server/pkg/logger"
	"sanguo-server/pkg/response"
)

// Recovery 捕获 panic，避免进程崩溃
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logger.WithFields(map[string]interface{}{
					"panic": r,
					"stack": string(debug.Stack()),
				}).Error("panic recovered")

				response.Fail(c, errors.CodeServer, "服务器开小差了",
					response.WithHTTPStatus(http.StatusInternalServerError))
				c.Abort()
			}
		}()
		c.Next()
	}
}
