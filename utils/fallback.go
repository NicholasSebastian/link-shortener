package utils

func Fallback[T comparable](v ...T) T {
	zeroVal := *new(T)
	for _, val := range v {
		if val != zeroVal {
			return val
		}
	}
	return zeroVal
}
