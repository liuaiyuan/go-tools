package array

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "T=NUMBERS,string"

// TIntersect 比较两个数组的值，并返回交集
// @param slice1 []T 与其他数组进行比较的第一个数组
// @param slice2 []T 与第一个数组进行比较的数组
// @return  value []T 与第一个数组进行比较的数组
// @date 2021-03-15 03:42:20
func TIntersect(slice1 []T, slice2 []T) (value []T) {
	for _, v := range slice1 {
		if TIn(v, slice2) {
			value = append(value, v)
		}
	}
	return
}
