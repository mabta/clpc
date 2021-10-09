package defs

import (
	"testing"
	"time"

	"github.com/mabta/clpc/eth/blocktime"
)

func TestDailyNSchedule(t *testing.T) {
	startHour := 7
	startMinute := 20
	duration := time.Minute * 5
	dailyTimes := 3
	diffDuration := 30 * time.Second

	ss := DailyNSchedule(startHour, startMinute, duration, diffDuration, dailyTimes)
	for _, s := range ss {
		t.Log(blocktime.TimeStr(s.Start), blocktime.TimeStr(s.End))
	}
}
