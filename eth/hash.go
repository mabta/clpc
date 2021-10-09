package eth

import "github.com/ethereum/go-ethereum/common"

// BlockHashToInt64 区块Hash的int64
func BlockHashToInt64(hash string) int64 {
	return common.HexToHash(hash).Big().Int64()
}

// BlockHashToUint64 区块Hash的uint64
func BlockHashToUint64(hash string) uint64 {
	return common.HexToHash(hash).Big().Uint64()
}
