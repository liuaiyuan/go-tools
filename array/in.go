package array

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "T=NUMBERS,string"

// TIn 搜索数组中是否存在指定的值。
func TIn(search T, slice []T) bool {
	for _, v := range slice {
		if v == search {
			return true
		}
	}
	return false
}
