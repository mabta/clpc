package lottery

import (
	"github.com/mabta/clpc/defs"
	"github.com/mabta/clpc/defs/draw"
)

type Ticket struct {
	*defs.Block
	Decimal int
	Drawer  draw.Drawer
}

func (t *Ticket) Draw() []draw.Result {
	return t.Drawer.Draw(t.Block, t.Decimal)
}
func (t *Ticket) SetBlock(block *defs.Block) {
	t.Block = block
}
