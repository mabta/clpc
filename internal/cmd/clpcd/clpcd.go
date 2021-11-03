package main

import (
	"log"

	"github.com/mabta/clpc/internal/cfg"
	"github.com/mabta/clpc/internal/db"
	LG "github.com/mabta/clpc/internal/log"
	"github.com/mabta/clpc/internal/lottery"
	"github.com/mabta/clpc/internal/lottery/handler"
	"github.com/mabta/clpc/internal/redis"
)

func init() {
	if err := cfg.Init(); err != nil {
		log.Fatal(err)
	}
	if err := LG.Init(cfg.Settings.Log); err != nil {
		log.Fatal(err)
	}
	if err := redis.InitFrom(cfg.Settings.Redis); err != nil {
		LG.Logger.Fatal(err)
	}
	if err := db.Init(cfg.Settings.DB.DSN); err != nil {
		LG.Logger.Fatal(err)
	}
}

func main() {
	for {
		LG.Logger.Info("服务器启动")
		engine, err := lottery.NewEngine(cfg.Settings.Eth.Provider, handler.DefaultHandler)
		if err != nil {
			LG.Logger.Error("连接以太坊失败：", err)
			continue
		}
		engine.LoadDescribies(cfg.Settings.Tickets)
		LG.Logger.Error(engine.Serve())
	}
}
