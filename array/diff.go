package array

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "T=NUMBERS,string"

// TDiff 比较两个数组的值，并返回差集
// @param slice1 []T 与其他数组进行比较的第一个数组
// @param slice2 []T 与第一个数组进行比较的数组
// @return  value []T 返回一个差集数组，该数组包括了所有在被比较的数组（slice1）中，但是不在任何其他参数数组（slice2）中的值。
// @date 2021-03-15 03:39:13
func TDiff(slice1 []T, slice2 []T) (value []T) {
	for _, v := range slice1 {
		if !TIn(v, slice2) {
			value = append(value, v)
		}
	}
	return
}
