package test

import (
	"github.com/liuaiyuan/go-tools/array"
	"testing"
)

func TestComb(t *testing.T) {
	var arr1 = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "12", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24"}
	var arr2 = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var arr3 = []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	var arr4 = []string{"1", "2", "3", "4", "5", "6", "7"}

	if v := array.StringComb(arr1, 5); len(v) != 42504 {
		t.Errorf("test number comb error")
	}
	if v := array.StringComb(arr1, 3); len(v) != 2024 {
		t.Errorf("test number comb error")
	}

	if v := array.StringComb(arr2, 5); len(v) != 126 {
		t.Errorf("test number comb error")
	}

	if v := array.StringComb(arr2, 3); len(v) != 84 {
		t.Errorf("test number comb error")
	}

	if v := array.StringComb(arr3, 5); len(v) != 56 {
		t.Errorf("test number comb error")
	}

	if v := array.StringComb(arr3, 3); len(v) != 56 {
		t.Errorf("test number comb error")
	}

	if v := array.StringComb(arr4, 5); len(v) != 21 {
		t.Errorf("test number comb error len:%d", len(v))
	}

	if v := array.StringComb(arr4, 3); len(v) != 35 {
		t.Errorf("test number comb error len:%d", len(v))
	}
}
