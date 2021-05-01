package str

import "strconv"

func ToFloat32D(s string, d float32) float32 {
	if v, err := strconv.ParseFloat(s, 32); err != nil {
		return d
	} else {
		return float32(v)
	}
}

func ToFloat64D(s string, d float64) float64 {
	if v, err := strconv.ParseFloat(s, 64); err != nil {
		return d
	} else {
		return v
	}
}

func ToIntD(s string, d int) int {
	if v, err := strconv.Atoi(s); err != nil {
		return d
	} else {
		return v
	}
}

func ToInt16D(s string, d int16) int16 {
	if v, err := strconv.ParseInt(s, 10, 16); err != nil {
		return d
	} else {
		return int16(v)
	}
}

func ToInt32D(s string, d int32) int32 {
	if v, err := strconv.ParseInt(s, 10, 32); err != nil {
		return d
	} else {
		return int32(v)
	}
}

func ToInt64D(s string, d int64) int64 {
	if v, err := strconv.ParseInt(s, 10, 64); err != nil {
		return d
	} else {
		return v
	}
}

func ToInt8D(s string, d int8) int8 {
	if v, err := strconv.ParseInt(s, 10, 8); err != nil {
		return d
	} else {
		return int8(v)
	}
}

func ToUintD(s string, d uint) uint {
	if v, err := strconv.ParseUint(s, 10, 64); err != nil {
		return d
	} else {
		return uint(v)
	}
}

func ToUint16D(s string, d uint16) uint16 {
	if v, err := strconv.ParseUint(s, 10, 16); err != nil {
		return d
	} else {
		return uint16(v)
	}
}

func ToUint32D(s string, d uint32) uint32 {
	if v, err := strconv.ParseUint(s, 10, 32); err != nil {
		return d
	} else {
		return uint32(v)
	}
}

func ToUint64D(s string, d uint64) uint64 {
	if v, err := strconv.ParseUint(s, 10, 64); err != nil {
		return d
	} else {
		return v
	}
}

func ToUint8D(s string, d uint8) uint8 {
	if v, err := strconv.ParseUint(s, 10, 8); err != nil {
		return 0
	} else {
		return uint8(v)
	}
}
