package util

// Must returns v if err is nil. If err is not nil, it panics.
//
// This is useful for static initialization of variables.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
