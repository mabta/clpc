package fns

import "time"

func DailySchedule(start time.Time, duration time.Duration, nums int) (schedules []uint64) {
	for i := 0; i < nums; i++ {
		item := uint64(start.Add(time.Duration(i) * duration).Unix())
		schedules = append(schedules, item)
	}
	return schedules
}

func DailyNSchedule(hour, minute int, duration time.Duration, nums int) (schedules []uint64) {
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, time.Local)
	return DailySchedule(start, duration, nums)
}
