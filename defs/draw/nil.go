package draw

import (
	"github.com/mabta/clpc/defs"
)

// Nil 无效的默认实现
type Nil struct{}

// Draw 开奖
func (n *Nil) Draw(block *defs.Block, decimal int) []Result {
	panic("请指定有效的算法实现")
}
