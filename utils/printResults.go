package utils

import "strconv"

func PrintResults(day int, results []string) {
	println("=------ Day " + strconv.Itoa(day) + " ------=")
	for i, result := range results {
		println("Part " + strconv.Itoa(i+1) + ": " + result)
	}
	println("=-------------------=")
}
