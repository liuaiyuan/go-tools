// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package array

// Float32Comb 求数组[]Float32无序组合
// @param arr []Float32 规定的数组
// @param num int 组合长度
// @return result [][]Float32 组合后的数组
// @date 2021-03-15 03:59:34
func Float32Comb(arr []float32, num int) (result [][]float32) {
	var arrLen = len(arr)

	if num <= 0 || num > arrLen {
		return result
	}

	for i := 0; i < arrLen; i++ {
		var tmp []float32
		tmp = append(tmp, arr[i])
		if num == 1 {
			result = append(result, tmp)
		} else {
			var c = Float32Comb(arr[i+1:], num-1)
			for j := 0; j < len(c); j++ {
				result = append(result, append(tmp, c[j]...))
			}
		}
	}

	return result
}

// Float64Comb 求数组[]Float64无序组合
// @param arr []Float64 规定的数组
// @param num int 组合长度
// @return result [][]Float64 组合后的数组
// @date 2021-03-15 03:59:34
func Float64Comb(arr []float64, num int) (result [][]float64) {
	var arrLen = len(arr)

	if num <= 0 || num > arrLen {
		return result
	}

	for i := 0; i < arrLen; i++ {
		var tmp []float64
		tmp = append(tmp, arr[i])
		if num == 1 {
			result = append(result, tmp)
		} else {
			var c = Float64Comb(arr[i+1:], num-1)
			for j := 0; j < len(c); j++ {
				result = append(result, append(tmp, c[j]...))
			}
		}
	}

	return result
}

// IntComb 求数组[]Int无序组合
// @param arr []Int 规定的数组
// @param num int 组合长度
// @return result [][]Int 组合后的数组
// @date 2021-03-15 03:59:34
func IntComb(arr []int, num int) (result [][]int) {
	var arrLen = len(arr)

	if num <= 0 || num > arrLen {
		return result
	}

	for i := 0; i < arrLen; i++ {
		var tmp []int
		tmp = append(tmp, arr[i])
		if num == 1 {
			result = append(result, tmp)
		} else {
			var c = IntComb(arr[i+1:], num-1)
			for j := 0; j < len(c); j++ {
				result = append(result, append(tmp, c[j]...))
			}
		}
	}

	return result
}

// Int16Comb 求数组[]Int16无序组合
// @param arr []Int16 规定的数组
// @param num int 组合长度
// @return result [][]Int16 组合后的数组
// @date 2021-03-15 03:59:34
func Int16Comb(arr []int16, num int) (result [][]int16) {
	var arrLen = len(arr)

	if num <= 0 || num > arrLen {
		return result
	}

	for i := 0; i < arrLen; i++ {
		var tmp []int16
		tmp = append(tmp, arr[i])
		if num == 1 {
			result = append(result, tmp)
		} else {
			var c = Int16Comb(arr[i+1:], num-1)
			for j := 0; j < len(c); j++ {
				result = append(result, append(tmp, c[j]...))
			}
		}
	}

	return result
}

// Int32Comb 求数组[]Int32无序组合
// @param arr []Int32 规定的数组
// @param num int 组合长度
// @return result [][]Int32 组合后的数组
// @date 2021-03-15 03:59:34
func Int32Comb(arr []int32, num int) (result [][]int32) {
	var arrLen = len(arr)

	if num <= 0 || num > arrLen {
		return result
	}

	for i := 0; i < arrLen; i++ {
		var tmp []int32
		tmp = append(tmp, arr[i])
		if num == 1 {
			result = append(result, tmp)
		} else {
			var c = Int32Comb(arr[i+1:], num-1)
			for j := 0; j < len(c); j++ {
				result = append(result, append(tmp, c[j]...))
			}
		}
	}

	return result
}

// Int64Comb 求数组[]Int64无序组合
// @param arr []Int64 规定的数组
// @param num int 组合长度
// @return result [][]Int64 组合后的数组
// @date 2021-03-15 03:59:34
func Int64Comb(arr []int64, num int) (result [][]int64) {
	var arrLen = len(arr)

	if num <= 0 || num > arrLen {
		return result
	}

	for i := 0; i < arrLen; i++ {
		var tmp []int64
		tmp = append(tmp, arr[i])
		if num == 1 {
			result = append(result, tmp)
		} else {
			var c = Int64Comb(arr[i+1:], num-1)
			for j := 0; j < len(c); j++ {
				result = append(result, append(tmp, c[j]...))
			}
		}
	}

	return result
}

// Int8Comb 求数组[]Int8无序组合
// @param arr []Int8 规定的数组
// @param num int 组合长度
// @return result [][]Int8 组合后的数组
// @date 2021-03-15 03:59:34
func Int8Comb(arr []int8, num int) (result [][]int8) {
	var arrLen = len(arr)

	if num <= 0 || num > arrLen {
		return result
	}

	for i := 0; i < arrLen; i++ {
		var tmp []int8
		tmp = append(tmp, arr[i])
		if num == 1 {
			result = append(result, tmp)
		} else {
			var c = Int8Comb(arr[i+1:], num-1)
			for j := 0; j < len(c); j++ {
				result = append(result, append(tmp, c[j]...))
			}
		}
	}

	return result
}

// UintComb 求数组[]Uint无序组合
// @param arr []Uint 规定的数组
// @param num int 组合长度
// @return result [][]Uint 组合后的数组
// @date 2021-03-15 03:59:34
func UintComb(arr []uint, num int) (result [][]uint) {
	var arrLen = len(arr)

	if num <= 0 || num > arrLen {
		return result
	}

	for i := 0; i < arrLen; i++ {
		var tmp []uint
		tmp = append(tmp, arr[i])
		if num == 1 {
			result = append(result, tmp)
		} else {
			var c = UintComb(arr[i+1:], num-1)
			for j := 0; j < len(c); j++ {
				result = append(result, append(tmp, c[j]...))
			}
		}
	}

	return result
}

// Uint16Comb 求数组[]Uint16无序组合
// @param arr []Uint16 规定的数组
// @param num int 组合长度
// @return result [][]Uint16 组合后的数组
// @date 2021-03-15 03:59:34
func Uint16Comb(arr []uint16, num int) (result [][]uint16) {
	var arrLen = len(arr)

	if num <= 0 || num > arrLen {
		return result
	}

	for i := 0; i < arrLen; i++ {
		var tmp []uint16
		tmp = append(tmp, arr[i])
		if num == 1 {
			result = append(result, tmp)
		} else {
			var c = Uint16Comb(arr[i+1:], num-1)
			for j := 0; j < len(c); j++ {
				result = append(result, append(tmp, c[j]...))
			}
		}
	}

	return result
}

// Uint32Comb 求数组[]Uint32无序组合
// @param arr []Uint32 规定的数组
// @param num int 组合长度
// @return result [][]Uint32 组合后的数组
// @date 2021-03-15 03:59:34
func Uint32Comb(arr []uint32, num int) (result [][]uint32) {
	var arrLen = len(arr)

	if num <= 0 || num > arrLen {
		return result
	}

	for i := 0; i < arrLen; i++ {
		var tmp []uint32
		tmp = append(tmp, arr[i])
		if num == 1 {
			result = append(result, tmp)
		} else {
			var c = Uint32Comb(arr[i+1:], num-1)
			for j := 0; j < len(c); j++ {
				result = append(result, append(tmp, c[j]...))
			}
		}
	}

	return result
}

// Uint64Comb 求数组[]Uint64无序组合
// @param arr []Uint64 规定的数组
// @param num int 组合长度
// @return result [][]Uint64 组合后的数组
// @date 2021-03-15 03:59:34
func Uint64Comb(arr []uint64, num int) (result [][]uint64) {
	var arrLen = len(arr)

	if num <= 0 || num > arrLen {
		return result
	}

	for i := 0; i < arrLen; i++ {
		var tmp []uint64
		tmp = append(tmp, arr[i])
		if num == 1 {
			result = append(result, tmp)
		} else {
			var c = Uint64Comb(arr[i+1:], num-1)
			for j := 0; j < len(c); j++ {
				result = append(result, append(tmp, c[j]...))
			}
		}
	}

	return result
}

// Uint8Comb 求数组[]Uint8无序组合
// @param arr []Uint8 规定的数组
// @param num int 组合长度
// @return result [][]Uint8 组合后的数组
// @date 2021-03-15 03:59:34
func Uint8Comb(arr []uint8, num int) (result [][]uint8) {
	var arrLen = len(arr)

	if num <= 0 || num > arrLen {
		return result
	}

	for i := 0; i < arrLen; i++ {
		var tmp []uint8
		tmp = append(tmp, arr[i])
		if num == 1 {
			result = append(result, tmp)
		} else {
			var c = Uint8Comb(arr[i+1:], num-1)
			for j := 0; j < len(c); j++ {
				result = append(result, append(tmp, c[j]...))
			}
		}
	}

	return result
}

// StringComb 求数组[]String无序组合
// @param arr []String 规定的数组
// @param num int 组合长度
// @return result [][]String 组合后的数组
// @date 2021-03-15 03:59:34
func StringComb(arr []string, num int) (result [][]string) {
	var arrLen = len(arr)

	if num <= 0 || num > arrLen {
		return result
	}

	for i := 0; i < arrLen; i++ {
		var tmp []string
		tmp = append(tmp, arr[i])
		if num == 1 {
			result = append(result, tmp)
		} else {
			var c = StringComb(arr[i+1:], num-1)
			for j := 0; j < len(c); j++ {
				result = append(result, append(tmp, c[j]...))
			}
		}
	}

	return result
}
