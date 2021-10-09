package eth

import (
	"testing"
	"time"

	"github.com/mabta/clpc/defs"
)

const testProvider = "wss://mainnet.infura.io/ws/v3/<YOUR KEY>"

func TestSubNewHeader(t *testing.T) {
	handler := func(block *defs.Block) {
		t.Log("新区块：", block.Hash)
		if block.Number < 0 {
			t.Fatal("区块高度错误")
		}
		if block.Time > uint64(time.Now().Unix()) {
			t.Fatal("区块时间异常")
		}
	}
	header, err := SubNewHeader(testProvider)
	if err != nil {
		t.Fatal(err)
	}
	header.SetHandler(handler)
	t.Fatal(header.Serve())
}
