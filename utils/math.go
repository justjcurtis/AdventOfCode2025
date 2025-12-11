package utils

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		a = -a
	}
	return a
}

func IntAbs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

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
	for range n {
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
