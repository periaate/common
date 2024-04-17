package common

import (
	"golang.org/x/exp/constraints"
)

type Numeric interface {
	~float32 | ~float64 | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func Abs[N Numeric](x N) (zero N) {
	if x < zero {
		return -x
	}
	return x
}

func Any[T any](arr []T, f func(T) bool) (ind int, ok bool) {
	for i, v := range arr {
		if f(v) {
			return i, true
		}
	}
	return -1, false
}

func Clamp[N Numeric](val, lower, upper N) (res N) {
	if val >= upper {
		return upper
	}
	if val <= lower {
		return lower
	}
	return val
}

func ZeroClamp[N Numeric](val, max N) N {
	if max < 0 {
		return Clamp(val, max, 0)
	}
	return Clamp(val, 0, max)
}

func SameSign[N Numeric](a, b N) bool {
	return (a > 0 && b > 0) || (a < 0 && b < 0)
}

func SmartClamp[I constraints.Integer](a, b I) I {
	switch {
	case b == 0 || a == 0:
		return 0
	case SameSign(a, b):
		return ZeroClamp(a, b)
	default:
		return (a % b) + b
	}
}
