package util

func Filter[T any](list []T, filterFunc func(v T) bool) []T {
	matches := 0
	for _, v := range list {
		if filterFunc(v) {
			matches++
		}
	}
	rv := make([]T, matches)
	i := 0
	for _, v := range list {
		if filterFunc(v) {
			rv[i] = v
			i++
		}
	}
	return rv
}
