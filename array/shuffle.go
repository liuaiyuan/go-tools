package array

import (
	"math/rand"
)

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "T=NUMBERS,string"

// TShuffle 打乱数组顺序
// @param slice []T 要打乱数组
// @return  []T  打乱后的数组
// @date 2021-03-15 03:52:19
func TShuffle(slice []T) []T {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
	return slice
}
