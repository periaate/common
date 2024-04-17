package common

func First[T any](arr []T, f func(T) bool) (res T, ind int) {
	for i, v := range arr {
		if f(v) {
			return v, i
		}
	}
	return res, -1
}

func Keys[T comparable, V any](m map[T]V) []T {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func All[T any, K comparable](mustBe K, fns ...func(T) K) func(T) K {
	return func(a T) K {
		for _, fn := range fns {
			if r := fn(a); r != mustBe {
				return r
			}
		}
		return mustBe
	}
}

func Negate[T any](fn func(T) bool) func(T) bool {
	return func(t T) bool {
		return !fn(t)
	}
}

func Pipe[T any](fns ...func(T) T) func(T) T {
	return func(t T) T {
		for _, fn := range fns {
			if fn != nil {
				t = fn(t)
			}
		}
		return t
	}
}
