package utils

func Clone2D[T any](src [][]T) [][]T {
	dst := make([][]T, len(src))
	for i := range src {
		row := make([]T, len(src[i]))
		copy(row, src[i])
		dst[i] = row
	}
	return dst
}
