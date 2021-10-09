package main

import (
	"log"

	"github.com/mabta/clpc/defs"
	"github.com/mabta/clpc/eth"
	"github.com/mabta/clpc/internal/cfg"
)

func init() {
	if err := cfg.Init(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	header, err := eth.SubNewHeader(cfg.Settings.Eth.Provider, func(block *defs.Block) {
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(header.Serve())
}
