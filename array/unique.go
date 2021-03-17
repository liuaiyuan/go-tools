package array

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "T=NUMBERS,string"

// TUnique 函数用于移除数组中重复的值
// @param slice []T 规定数组
// @return  value []T 唯一数组
// @date 2021-03-15 03:46:38
func TUnique(slice []T) (value []T) {
	for _, v := range slice {
		if !TIn(v, value) {
			value = append(value, v)
		}
	}
	return
}
