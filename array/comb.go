package array

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "T=NUMBERS,string"

// TComb 求数组[]T无序组合
func TComb(arr []T, num int) (result [][]T) {
	var arrLen = len(arr)

	if num <= 0 || num > arrLen {
		return result
	}

	for i := 0; i < arrLen; i++ {
		var tmp []T
		tmp = append(tmp, arr[i])
		if num == 1 {
			result = append(result, tmp)
		} else {
			var c = TComb(arr[i+1:], num-1)
			for j := 0; j < len(c); j++ {
				result = append(result, append(tmp, c[j]...))
			}
		}
	}

	return result
}
