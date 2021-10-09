package eth

import (
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mabta/clpc/defs"
)

// NewHeaderSubHandler 新块的处理函数
type NewHeaderSubHandler func(block *defs.Block)

// Header 以太坊新块
type Header struct {
	ch      chan *types.Header
	sub     ethereum.Subscription
	handler NewHeaderSubHandler
	client  *ethclient.Client
}

// SubNewHeader 订阅以太坊新块
func SubNewHeader(provider string) (*Header, error) {
	client, err := ethclient.Dial(provider)
	if err != nil {
		return nil, err
	}

	h := &Header{
		ch:     make(chan *types.Header),
		client: client,
	}

	sub, err := client.SubscribeNewHead(context.Background(), h.ch)
	if err != nil {
		return nil, err
	}
	h.sub = sub
	return h, nil
}
func (h *Header) SetHandler(handler NewHeaderSubHandler) {
	h.handler = handler
}

// Serve 处理以太坊新块
func (h *Header) Serve() error {
	for {
		select {
		case err := <-h.sub.Err():
			h.client.Close()
			return err
		case header := <-h.ch:
			h.handler(defs.NewHeaderBlock(header))
		}
	}
}
