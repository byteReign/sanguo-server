package redis

import (
	"context"
	"fmt"
	"sync"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

var (
	client *goredis.Client
	once   sync.Once
)

// Config Redis 配置
type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	PoolSize int    `yaml:"poolSize"`
}

// Init 初始化 Redis 连接（单例，仅首次调用生效）
func Init(cfg Config) error {
	var err error
	once.Do(func() {
		err = initRedis(cfg)
	})
	return err
}

func initRedis(cfg Config) error {
	opts := &goredis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	}
	if cfg.PoolSize > 0 {
		opts.PoolSize = cfg.PoolSize
	}

	instance := goredis.NewClient(opts)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := instance.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("ping redis: %w", err)
	}

	client = instance
	return nil
}

// GetRedis 获取 Redis 单例
func GetRedis() *goredis.Client {
	return client
}
