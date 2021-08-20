package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pigeonligh/easygo/elog"
	"github.com/pigeonligh/ffxiv-todo/pkg/cache"
)

type Config struct {
	Debug bool   `name:"debug"`
	Port  int    `name:"port"`
	Data  string `name:"data"`
}

func RunServer(config *Config) error {
	if config.Debug {
		gin.SetMode(gin.DebugMode)
		elog.Debug()
	} else {
		gin.SetMode(gin.ReleaseMode)
		elog.Default()
	}

	addr := fmt.Sprintf(":%d", config.Port)
	manager := cache.New(config.Data)
	cache.Init(manager)
	elog.Info("Inited")

	router := gin.Default()

	Install(router)

	if err := router.Run(addr); err != nil {
		elog.Fatal(err)
	}
	return nil
}
