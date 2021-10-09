package draw

import "github.com/mabta/clpc/defs"

// Result 开奖结果
type Result int

// Drawer 开奖算法
type Drawer interface {
	Draw(*defs.Block, int) []Result
}

// inResult 结果是否在列表中
func inResult(r Result, rs []Result) bool {
	for _, item := range rs {
		if item == r {
			return true
		}
	}
	return false
}
