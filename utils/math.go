package utils

func LengthOfInt(n int) int {
	length := 0
	for n > 0 {
		n /= 10
		length++
	}
	return length
}

func IntPow10(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}

func IntMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func IntMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func GetSign(x int) int {
	if x < 0 {
		return -1
	}
	return 1
}
