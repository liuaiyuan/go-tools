package str

import "strconv"

func ToFloat32(s string) float32 {
	if v, err := strconv.ParseFloat(s, 32); err != nil {
		return 0
	} else {
		return float32(v)
	}
}

func ToFloat64(s string) float64 {
	if v, err := strconv.ParseFloat(s, 64); err != nil {
		return 0
	} else {
		return v
	}
}

func ToInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func ToInt16(s string) int16 {
	if v, err := strconv.ParseInt(s, 10, 16); err != nil {
		return 0
	} else {
		return int16(v)
	}
}

func ToInt32(s string) int32 {
	if v, err := strconv.ParseInt(s, 10, 32); err != nil {
		return 0
	} else {
		return int32(v)
	}
}

func ToInt64(s string) int64 {
	v, _ := strconv.ParseInt(s, 10, 64)
	return v
}

func ToInt8(s string) int8 {
	if v, err := strconv.ParseInt(s, 10, 8); err != nil {
		return 0
	} else {
		return int8(v)
	}
}

func ToUint(s string) uint {
	if v, err := strconv.ParseUint(s, 10, 64); err != nil {
		return 0
	} else {
		return uint(v)
	}
}

func ToUint16(s string) uint16 {
	if v, err := strconv.ParseUint(s, 10, 16); err != nil {
		return 0
	} else {
		return uint16(v)
	}
}

func ToUint32(s string) uint32 {
	if v, err := strconv.ParseUint(s, 10, 32); err != nil {
		return 0
	} else {
		return uint32(v)
	}
}

func ToUint64(s string) uint64 {
	v, _ := strconv.ParseUint(s, 10, 64)
	return v
}

func ToUint8(s string) uint8 {
	if v, err := strconv.ParseUint(s, 10, 8); err != nil {
		return 0
	} else {
		return uint8(v)
	}
}

