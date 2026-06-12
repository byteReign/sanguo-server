package router

import (
	"sanguo-server/internal/user"
	"sanguo-server/pkg/logger"

	"github.com/gin-gonic/gin"
)

func Init(g *gin.Engine) {
	g.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "sanguo-server is running"})
	})
	logger.Infof("user router initialized")
	user.Router(g)
}
