package test

import (
	"github.com/liuaiyuan/go-tools/array"
	"testing"
)

func TestSum(t *testing.T) {
	var arr1 = []int64{10, 100, 1000}

	if sum := array.Int64Sum(arr1); sum != 1110 {
		t.Errorf("array sum error: %d", sum)
	}

	t.Logf("array sum test success")
}
