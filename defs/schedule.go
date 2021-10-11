package defs

import (
	"fmt"
	"time"

	"github.com/mabta/clpc/eth/blocktime"
)

type Schedule struct {
	Start, End uint64
}

// IsIn 判断指定时间是否在计划时间内
func (s *Schedule) IsIn(t uint64) bool {
	return t >= s.Start && t <= s.End
}

// String 将计划时间用易于阅读的形式呈现
func (s *Schedule) String() string {
	return fmt.Sprintf("%s - %s", blocktime.TimeStr(s.Start), blocktime.TimeStr(s.End))
}

// FullString 将计划时间用易于阅读的形式呈现(包含日期)
func (s *Schedule) FullString() string {
	return fmt.Sprintf("%s - %s", blocktime.DateTimeStr(s.Start), blocktime.DateTimeStr(s.End))
}
func (s *Schedule) IssueStr() string {
	return blocktime.IssueStr(s.Start)
}

// DailySchedule 生成当天计划
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

// DailyNSchedule 通过指定整数形式的时间来生成当天计划
func DailyNSchedule(hour, minute int, duration, diffDuration time.Duration, nums int) (schedules []*Schedule) {
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, time.Local)
	return DailySchedule(start, duration, diffDuration, nums)
}
