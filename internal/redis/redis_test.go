package redis

import (
	"testing"

	"github.com/mabta/clpc/internal/cfg"
)

func testInit(t *testing.T) {
	if err := cfg.InitFrom("../../config.json"); err != nil {
		t.Fatal(err)
	}
	if err := InitFrom(cfg.Settings.Redis); err != nil {
		t.Fatal(err)
	}
}

func TestInitFrom(t *testing.T) {
	testInit(t)
}

func TestExists(t *testing.T) {
	testInit(t)
	key := "foo"
	if err := Set(key, "bar"); err != nil {
		t.Fatal(err)
	}
	exists, err := Exists("foo")
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatal(key, "必须存在")
	}
}
func TestSet(t *testing.T) {
	testInit(t)
	if err := Set("foo", "bar"); err != nil {
		t.Fatal(err)
	}
}

func TestGet(t *testing.T) {
	testInit(t)
	if err := Set("foo", "bar"); err != nil {
		t.Fatal(err)
	}
	val, err := Get("foo")
	if err != nil {
		t.Fatal(err)
	}
	if val != "bar" {
		t.Fatal("必须是bar")
	}
}
