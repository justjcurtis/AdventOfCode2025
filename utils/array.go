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

func ConvertToByteMatrix(input []string) [][]byte {
	matrix := make([][]byte, len(input))
	width := len(input[0])
	for i := range input {
		matrix[i] = make([]byte, width)
		for j := range input[i] {
			matrix[i][j] = input[i][j]
		}
	}
	return matrix
}
