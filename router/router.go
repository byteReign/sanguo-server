package router

import (
	"sanguo-server/internal/user"
	"sanguo-server/pkg/logger"

	"github.com/gin-gonic/gin"
)

func Init(g *gin.Engine) {
	logger.Infof("user router initialized")
	user.Router(g)
}
