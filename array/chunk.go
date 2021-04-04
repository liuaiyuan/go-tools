package array

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "T=NUMBERS,string"

// TChunk 数组分割为新的数组块
func TChunk(items []T, size int) (chunks [][]T) {

	for size < len(items) {
		items, chunks = items[size:], append(chunks, items[0:size:size])
	}

	return append(chunks, items)

}
