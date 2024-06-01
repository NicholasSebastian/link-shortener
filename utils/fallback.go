package utils

// Given a list of arguments, returns the first argument that is not of zero value.

func Fallback[T comparable](v ...T) T {
	zeroVal := *new(T)
	for _, val := range v {
		if val != zeroVal {
			return val
		}
	}
	return zeroVal
}
