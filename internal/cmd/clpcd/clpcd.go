package main

import (
	"log"

	"github.com/mabta/clpc/internal/cfg"
	"github.com/mabta/clpc/internal/lottery"
	"github.com/mabta/clpc/internal/lottery/handler"
)

func init() {
	if err := cfg.InitFrom("../../../config.json"); err != nil {
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
