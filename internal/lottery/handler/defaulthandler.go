package handler

import (
	"fmt"

	"github.com/mabta/clpc/defs"
	"github.com/mabta/clpc/internal/lottery"
)

func DefaultHandler(describies []*lottery.Describe, block *defs.Block) {
	fmt.Println(block.Number)
	for _, d := range describies {
		fmt.Println(d.Name, d.Schedule())
	}
}
