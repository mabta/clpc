package handler

import (
	"fmt"

	"github.com/mabta/clpc/defs"
	"github.com/mabta/clpc/internal/lottery"
)

func DefaultHandler(describies []*lottery.Describe, block *defs.Block) {
	fmt.Println("aaa")
	fmt.Println(describies)
	fmt.Println(block.Number)
}
