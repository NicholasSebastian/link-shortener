package utils

func Fallback(v ...string) string {
	for _, val := range v {
		if val != "" {
			return val
		}
	}
	return ""
}
