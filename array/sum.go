package array

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "N=NUMBERS"

// TSum T数组求和
func NSum(arr []N) (v N) {
	for _, i := range arr {
		v += i
	}
	return v
}
