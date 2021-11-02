package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mabta/clpc/internal/api/v1/handler"
	"github.com/mabta/clpc/internal/cfg"
	"github.com/mabta/clpc/internal/db"
)

func init() {
	if err := cfg.Init(); err != nil {
		log.Fatal(err)
	}
	if err := db.Init(cfg.Settings.DB.DSN); err != nil {
		log.Fatal(err)
	}
}

func main() {
	e := gin.New()
	e.Use(gin.Logger(), gin.Recovery())
	handler.RegisterRouter(e)
	log.Fatal(e.Run(cfg.Settings.Web.Addr))
}
