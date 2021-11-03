package cfg

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Config struct {
	Eth     *EthConfig      `json:"eth"`
	Web     *WebConfig      `json:"web"`
	Tickets []*TicketConfig `json:"tickets"`
	Redis   *RedisConfig    `json:"redis"`
	DB      *DBConfig       `json:"db"`
	Log     *LogConfig      `json:"log"`
}

type EthConfig struct {
	Provider string `json:"provider"`
}

type WebConfig struct {
	Addr string `json:"addr"`
}

type TicketConfig struct {
	Name         string        `json:"name"`
	StartHour    int           `json:"start_hour"`
	StartMinute  int           `json:"start_minute"`
	Duration     time.Duration `json:"duration"`
	DiffDuration time.Duration `json:"diff_duration"`
	Decimal      int           `json:"decimal"`
	DailyTimes   int           `json:"daily_times"`
	Drawer       string        `json:"drawer"`
}
type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}
type DBConfig struct {
	DSN string `json:"dsn"`
}
type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"max_size"`
	MaxBackups int    `json:"max_backups"`
	MaxAge     int    `json:"max_age"`
	Compress   bool   `json:"compress"`
}

var Settings *Config

func InitFrom(filename string) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	Settings = new(Config)
	if err := json.Unmarshal(buf, Settings); err != nil {
		return err
	}
	for _, t := range Settings.Tickets {
		t.Duration = t.Duration * time.Minute
		t.DiffDuration = t.DiffDuration * time.Second
	}
	if Settings.Log.Filename == "" {
		Settings.Log.Filename = "./clpc.log"
	}
	return nil
}

func Init() error {
	return InitFrom("./config.json")
}
