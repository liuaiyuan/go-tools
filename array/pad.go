package array

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "T=NUMBERS,string"

// TPad 将指定数量的带有指定值的元素插入到数组中。
// @param slice []T 规定数组
// @param size int 规定从函数返回的数组元素个数
// @param val T 规定从函数返回的数组中新元素的值
// @return  []T 返回后的数组
// @date 2021-03-15 03:54:47
func TPad(slice []T, size int, val T) []T {
	if size <= len(slice) {
		return slice
	}
	for i := 0; i < (size - len(slice)); i++ {
		slice = append(slice, val)
	}
	return slice
}
