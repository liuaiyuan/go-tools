package test

import (
	"github.com/liuaiyuan/go-tools/array"
	"testing"
)

func TestIntersect(t *testing.T) {
	if v := array.StringIntersect([]string{"A", "B", "C"}, []string{"C", "D", "A", "E"}); len(v) != 2 {
		t.Errorf("test string arr intersect error:%v", v)
	}

	if v := array.StringIntersect([]string{"A", "B", "C", "E"}, []string{"C", "D", "A", "E", "X", "Q"}); len(v) != 3 {
		t.Errorf("test string arr intersect error:%v", v)
	}

	if v := array.StringIntersect([]string{"X", "Y", "Z"}, []string{"C", "D", "A", "E"}); len(v) != 0 {
		t.Errorf("test string arr intersect error:%v", v)
	}

	t.Logf("intersect string array test success")
}
