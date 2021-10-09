package cfg

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Config struct {
	Eth     *EthConfig      `json:"eth"`
	Tickets []*TicketConfig `json:"tickets"`
}

type EthConfig struct {
	Provider string `json:"provider"`
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
	return nil
}

func Init() error {
	return InitFrom("./config.json")
}
