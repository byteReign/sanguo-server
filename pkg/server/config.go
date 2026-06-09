package server

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

// Config HTTP 服务配置
type Config struct {
	Host string `yaml:"host"` // 监听地址，空或 0.0.0.0 表示所有网卡
	Port int    `yaml:"port"` // 监听端口
	Mode string `yaml:"mode"` // debug / release / test
}

// Addr 返回 gin 监听地址，如 :8088
func (c Config) Addr() string {
	port := c.Port
	if port == 0 {
		port = 8080
	}

	host := strings.TrimSpace(c.Host)
	if host == "" || host == "0.0.0.0" {
		return fmt.Sprintf(":%d", port)
	}
	return fmt.Sprintf("%s:%d", host, port)
}

// ApplyMode 设置 gin 运行模式
func (c Config) ApplyMode() {
	switch strings.ToLower(c.Mode) {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}
