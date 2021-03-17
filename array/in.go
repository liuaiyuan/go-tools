package array

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "T=NUMBERS,string"

// TIn 搜索数组中是否存在指定的值。
// @param search T 要索索的值
// @param slice []T 要搜索的数组
// @return  bool 是否存在
// @date 2021-03-15 03:35:25
func TIn(search T, slice []T) bool {
	for _, v := range slice {
		if v == search {
			return true
		}
	}
	return false
}
