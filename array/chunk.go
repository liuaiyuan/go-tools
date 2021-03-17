package array

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "T=NUMBERS,string"

// TChunk 数组分割为新的数组块
// @param slice []T 规定要使用的数组
// @param size int 规定每个新数组块包含多少个元素
// @return  v [][]T 分割的数组
// @date 2021-03-15 03:49:44
func TChunk(items []T, size int) (chunks [][]T) {

	for size < len(items) {
		items, chunks = items[size:], append(chunks, items[0:size:size])
	}

	return append(chunks, items)

}
