package common

import "sort"

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

func Any[T any](ifAny bool, fns ...func(T) bool) func(T) bool {
	return func(a T) bool {
		for _, fn := range fns {
			if fn(a) == ifAny {
				return ifAny
			}
		}
		return !ifAny
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

func Collect[T any, K any](arr []T, fn func(T) K) []K {
	res := make([]K, 0, len(arr))
	for _, v := range arr {
		res = append(res, fn(v))
	}
	return res
}

// SplitWithAll splits given string into an array, using all other `match` strings as
// delimiters. String is matched using the longest delimiter first.
// If no match strings are given, the original string is returned.
// If no matches are found, the original string is returned.
// Matched delimiters are not included in the result.
// If a found match would add a zero-length string to the result, it is ignored.
// Any consecutive matches are treated as one.
// If an empty match string is given (i.e. ""), every character is split.
// Keep argument determines whether the split part is included in the split part.
func SplitWithAll(str string, keep bool, match ...string) (res []string) {
	mult := 1
	if keep {
		mult = 0
	}

	if len(match) == 0 || len(str) == 0 {
		return []string{str}
	}

	sort.SliceStable(match, func(i, j int) bool {
		return len(match[i]) > len(match[j])
	})

	var lastI int

	for i := 0; i < len(str); i++ {
		for _, pattern := range match {
			if i+len(pattern) > len(str) {
				continue
			}

			if str[i:i+len(pattern)] != pattern {
				continue
			}

			if len(str[lastI:i]) != 0 {
				res = append(res, str[lastI:i])
			}

			lastI = i + len(pattern)*mult
			if len(pattern) != 0 {
				i += len(pattern) - 1
			}
			break
		}
	}

	if len(str[lastI:]) != 0 {
		res = append(res, str[lastI:])
	}

	if len(res) == 0 {
		return []string{str}
	}

	return res
}

func Combinations[T any](a []T, B []T) (res [][2]T) {
	res = make([][2]T, 0, len(a)*len(B))
	for _, v := range a {
		for _, w := range B {
			res = append(res, [2]T{v, w})
		}
	}

	return res
}

func Join[T any](a ...[]T) (res []T) {
	for _, arr := range a {
		res = append(res, arr...)
	}
	return res
}
