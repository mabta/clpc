package lottery

import (
	"github.com/mabta/clpc/defs"
	"github.com/mabta/clpc/eth"
	"github.com/mabta/clpc/internal/cfg"
)

type EngineHandler func([]*Describe, *defs.Block)
type Engine struct {
	describies []*Describe
	head       *eth.Header
	handler    EngineHandler
}

func NewEngine(provider string, handler EngineHandler) (*Engine, error) {
	headSub, err := eth.SubNewHeader(provider)
	if err != nil {
		return nil, err
	}
	return &Engine{
		head:       headSub,
		handler:    handler,
		describies: make([]*Describe, 0),
	}, err
}

func (e *Engine) AddDescribe(t *Describe) *Engine {
	e.describies = append(e.describies, t)
	return e
}
func (e *Engine) LoadDescribies(cfgTickets []*cfg.TicketConfig) {
	for _, t := range cfgTickets {
		d := LoadDescribe(t)
		e.AddDescribe(d)
	}
}

func (e *Engine) Serve() error {
	e.head.SetHandler(func(block *defs.Block) {
		e.handler(e.describies, block)
	})
	if err := e.head.Serve(); err != nil {
		return err
	}
	return nil
}
