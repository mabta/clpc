package defs

import (
	"time"

	"github.com/mabta/clpc/eth/blocktime"
)

type Schedule struct {
	Start, End uint64
}

func DailySchedule(start time.Time, duration, diffDuration time.Duration, nums int) (schedules []*Schedule) {
	for i := 0; i < nums; i++ {
		s := &Schedule{}
		s.Start = uint64(start.Add(time.Duration(i) * duration).Unix())
		schedules = append(schedules, s)
	}
	for _, s := range schedules {
		s.End = uint64(blocktime.ToTime(s.Start).Add(duration).Add(-diffDuration).Unix())
	}
	return schedules
}

func DailyNSchedule(hour, minute int, duration, diffDuration time.Duration, nums int) (schedules []*Schedule) {
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, time.Local)
	return DailySchedule(start, duration, diffDuration, nums)
}
