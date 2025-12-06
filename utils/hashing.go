package utils

import "math"

func SzudzikPairing(x int, y int) int {
	if x < y {
		return y*y + x
	}
	return x*x + x + y
}

func SzudzikUnpairing(z int) (int, int) {
	x := int(math.Sqrt(float64(z)))
	if x*x == z {
		return x, 0
	}
	y := z - x*x
	if y < x {
		return y, x
	}
	return x, y - x
}

func TwoDToOneD(x int, y int, width int) int {
	return y*width + x
}

func OneDTwoD(z int, width int) (int, int) {
	return z % width, z / width
}
