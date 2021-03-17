package array

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "N=NUMBERS"

// TSum T数组求和
// @param arr []T 数组
// @return  v T 和值
// @date 2021-03-15 03:24:25
func NSum(arr []N) (v N) {
	for _, i := range arr {
		v += i
	}
	return v
}
