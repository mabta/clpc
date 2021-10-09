package draw

import (
	"math/rand"

	"github.com/mabta/clpc/defs"
	"github.com/mabta/clpc/eth"
)

// Simple 简单随机算法
type Simple struct{}

// Draw 开奖
func (s *Simple) Draw(block *defs.Block, decimal int) []Result {
	rand.Seed(eth.BlockHashToInt64(block.Hash))
	result := make([]Result, decimal)
	for i := 0; i < decimal; i++ {
		for {
			r := Result((rand.Int() % decimal) + 1)
			if !inResult(r, result) {
				result[i] = r
				break
			}
		}
	}
	return result
}
