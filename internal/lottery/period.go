package lottery

import (
	"fmt"
	"strconv"

	"github.com/mabta/clpc/defs"
	"github.com/mabta/clpc/eth/blocktime"
)

// GetPeriod 获取期数
func GetPeriod(schedule uint64, idx int) uint64 {
	dateStr := blocktime.DateStr(schedule)
	s := fmt.Sprintf("%s%03d", dateStr, idx+1)
	i, _ := strconv.ParseUint(s, 10, 64)
	return i
}

// GetNextPeriod 获取下一期
func GetNextPeriod(idx int, schedules []*defs.Schedule) (nextPeroid, nextPeroidSchedule uint64, ok bool) {
	schedulesNum := len(schedules)
	lastIdx := schedulesNum - 1
	nextIdx := idx + 1
	if nextIdx > lastIdx {
		return 0, 0, false
	}
	nextSchedule := schedules[nextIdx].Start
	return GetPeriod(nextSchedule, nextIdx), nextSchedule, true
}
