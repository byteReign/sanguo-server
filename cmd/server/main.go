package main

import (
	"flag"
	"fmt"
	"sanguo-server/pkg/config"
	"sanguo-server/pkg/database"
	"sanguo-server/pkg/logger"
	"sanguo-server/pkg/redis"
	"sanguo-server/router"

	"github.com/gin-gonic/gin"
)

func main() {
	configPath := flag.String("config", "configs/dev.yaml", "配置文件路径")
	flag.Parse()

	conf, err := config.Load(*configPath)
	if err != nil {
		panic(fmt.Sprintf("load config fail:%v", err))
	}
	if conf == nil {
		panic(fmt.Sprintf("load config is error"))
	}

	if err = logger.Init(conf.Log); err != nil {
		panic(fmt.Sprintf("init logger fail:%v", err))
	}

	if err = database.Init(conf.Database); err != nil {
		panic(fmt.Sprintf("init database fail:%v", err))
	}

	if err = redis.Init(conf.Redis); err != nil {
		panic(fmt.Sprintf("init redis fail:%v", err))
	}

	conf.Server.ApplyMode()

	engine := gin.Default()
	router.Init(engine)

	addr := conf.Server.Addr()
	logger.Infof("server starting on %s", addr)
	if err = engine.Run(addr); err != nil {
		panic(fmt.Sprintf("server run fail:%v", err))
	}
}
