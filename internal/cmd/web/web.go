package main

import (
	SYSLOG "log"

	"github.com/gin-gonic/gin"
	"github.com/mabta/clpc/internal/api/v1/handler"
	"github.com/mabta/clpc/internal/cfg"
	"github.com/mabta/clpc/internal/db"
	"github.com/mabta/clpc/internal/log"
)

func init() {
	if err := cfg.Init(); err != nil {
		SYSLOG.Fatal(err)
	}
	if err := log.Init(cfg.Settings.Log); err != nil {
		SYSLOG.Fatal(err)
	}
	if err := db.Init(cfg.Settings.DB.DSN); err != nil {
		log.Logger.Fatal(err)
	}
}

func main() {
	e := gin.New()
	e.Use(log.GinLogger(log.Logger.Desugar()))
	e.Use(log.GinRecovery(log.Logger.Desugar(), true))
	handler.RegisterRouter(e)
	log.Logger.Fatal(e.Run(cfg.Settings.Web.Addr))
}
