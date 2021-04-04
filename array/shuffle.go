package array

import (
	"math/rand"
)

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "T=NUMBERS,string"

// TShuffle 打乱数组顺序
func TShuffle(slice []T) []T {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
	return slice
}
