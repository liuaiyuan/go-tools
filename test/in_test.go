package test

import (
	"github.com/liuaiyuan/go-tools/array"
	"testing"
)

func TestIn(t *testing.T) {
	var arr = []string{"A", "B", "C", "D", "E", "F", "G"}

	if ok := array.StringIn("A", arr); !ok {
		t.Errorf("inString error")
	}

	if ok := array.StringIn("H", arr); ok {
		t.Errorf("inString error")
	}

	t.Logf("in string array test success")
}

