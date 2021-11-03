package main

import (
	"log"

	"github.com/mabta/clpc/internal/cfg"
	"github.com/mabta/clpc/internal/db"
	"github.com/mabta/clpc/internal/lottery"
	"github.com/mabta/clpc/internal/lottery/handler"
	"github.com/mabta/clpc/internal/redis"
)

func init() {
	if err := cfg.Init(); err != nil {
		log.Fatal(err)
	}
	if err := redis.InitFrom(cfg.Settings.Redis); err != nil {
		log.Fatal(err)
	}
	if err := db.Init(cfg.Settings.DB.DSN); err != nil {
		log.Fatal(err)
	}
}

func main() {
	for {
		log.Println("启动服务端")
		engine, err := lottery.NewEngine(cfg.Settings.Eth.Provider, handler.DefaultHandler)
		if err != nil {
			log.Println("连接以太坊网络出错：", err)
			continue
		}
		engine.LoadDescribies(cfg.Settings.Tickets)
		log.Println(engine.Serve())
	}
}
