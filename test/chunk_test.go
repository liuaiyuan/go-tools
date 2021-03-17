package test

import (
	"github.com/liuaiyuan/go-tools/array"
	"testing"
)

func TestChunk(t *testing.T) {
	var arr = []string{"A", "B", "C", "D", "E", "F", "G"}

	if chunks := array.StringChunk(arr, 1); len(chunks) != 7 {
		t.Errorf("string chunk error len != 1")
	}

	if chunks := array.StringChunk(arr, 2); len(chunks) != 4 {
		t.Errorf("string chunk error len != 3")
	}

	if chunks := array.StringChunk(arr, 3); len(chunks) != 3 {
		t.Errorf("string chunk error len != 3")
	}

	if chunks := array.StringChunk(arr, 5); len(chunks) != 2 {
		t.Errorf("string chunk error len != 5")
	}

	t.Logf("string array chunk test success")
}

