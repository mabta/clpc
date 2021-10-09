package cfg

import (
	"strings"
	"testing"
)

func TestInitFrom(t *testing.T) {
	if err := InitFrom("../../config.json"); err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(Settings.Eth.Provider, "wss://") {
		t.Fatal("eth的提供者需要wss协议")
	}
}
