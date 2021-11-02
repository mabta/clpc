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
	engine, err := lottery.NewEngine(cfg.Settings.Eth.Provider, handler.DefaultHandler)
	if err != nil {
		log.Fatal(err)
	}
	engine.LoadDescribies(cfg.Settings.Tickets)
	log.Fatal(engine.Serve())
}
