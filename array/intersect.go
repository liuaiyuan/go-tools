package array

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "T=NUMBERS,string"

// TIntersect 比较两个数组的值，并返回交集
func TIntersect(slice1 []T, slice2 []T) (value []T) {
	for _, v := range slice1 {
		if TIn(v, slice2) {
			value = append(value, v)
		}
	}
	return
}
