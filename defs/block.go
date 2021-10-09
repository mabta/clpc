package defs

import (
	"github.com/ethereum/go-ethereum/core/types"
)

// Block 区块
type Block struct {
	Hash   string
	Number uint64
	Time   uint64
}

// NewHeaderBlock 从以太坊新块构建区块对象
func NewHeaderBlock(h *types.Header) *Block {
	return &Block{
		Hash:   h.Hash().Hex(),
		Number: h.Number.Uint64(),
		Time:   h.Time,
	}
}
