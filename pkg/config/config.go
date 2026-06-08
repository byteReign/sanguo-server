package config

import (
	"fmt"
	"os"

	"sanguo-server/pkg/database"
	"sanguo-server/pkg/logger"
	"sanguo-server/pkg/redis"

	"gopkg.in/yaml.v3"
)

// Config 应用配置
type Config struct {
	Log      logger.Config    `yaml:"log"`
	Database database.Config  `yaml:"database"`
	Redis    redis.Config     `yaml:"redis"`
}

// Load 从 YAML 文件加载配置
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse config file: %w", err)
	}

	return &cfg, nil
}
