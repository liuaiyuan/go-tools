package test

import (
	"github.com/liuaiyuan/go-tools/array"
	"testing"
)

func TestDiff(t *testing.T) {

	if v := array.StringDiff([]string{"1", "2"}, []string{"2", "3"}); !array.StringIn("1", v) && len(v) != 1 {
		t.Errorf("test string diff error:%v", v)
	}

	if v := array.StringDiff([]string{"1", "2", "3", "4"}, []string{"2", "3"}); len(v) != 2 {
		t.Errorf("test string diff error:%v", v)
	}
}
