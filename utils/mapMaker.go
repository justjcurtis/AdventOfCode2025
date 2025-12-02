package utils

func CountMap[T comparable](arr []T) map[T]int {
	counts := make(map[T]int)
	for _, item := range arr {
		counts[item]++
	}
	return counts
}
