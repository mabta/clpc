package lottery

import (
	"time"

	"github.com/mabta/clpc/defs/draw"
	"github.com/mabta/clpc/defs/lottery"
	"github.com/mabta/clpc/internal/cfg"
)

type Describe struct {
	*lottery.Ticket
	Name         string
	StartHour    int
	StartMinute  int
	Duration     time.Duration
	DiffDuration time.Duration
	DailyTimes   int
}

func NewDescribe(name string, decimal, startHour, startMinute, dailyTimes int, duration, diffDuration time.Duration, drawer draw.Drawer) *Describe {
	return &Describe{
		Ticket: &lottery.Ticket{
			Decimal: decimal,
			Drawer:  drawer,
		},
		Name:         name,
		StartHour:    startHour,
		StartMinute:  startMinute,
		Duration:     duration,
		DiffDuration: diffDuration,
		DailyTimes:   dailyTimes,
	}
}

func LoadDescribe(ct *cfg.TicketConfig) *Describe {
	return NewDescribe(ct.Name, ct.Decimal, ct.StartHour, ct.StartMinute, ct.DailyTimes, ct.Duration, ct.DiffDuration, getDrawer(ct.Drawer))
}

func getDrawer(drawerName string) draw.Drawer {
	switch drawerName {
	case "simple":
		return &draw.Simple{}
	}
	return &draw.Nil{}
}
