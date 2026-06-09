package user

import "github.com/gin-gonic/gin"

func Router(g *gin.Engine) {
	user := g.Group("user")

	user.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello world"})
	})
}
