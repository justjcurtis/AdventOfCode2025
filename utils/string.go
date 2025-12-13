package utils

import "slices"

func Atoi(s string) int {
	result := 0
	for _, char := range s {
		result = result*10 + int(char-'0')
	}
	return result
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	isNegative := n < 0
	buffer := []rune{}
	for n != 0 {
		digit := IntAbs(n % 10)
		buffer = append(buffer, rune('0'+digit))
		n /= 10
	}
	if isNegative {
		buffer = append(buffer, '-')
	}
	slices.Reverse(buffer)
	return string(buffer)
}
