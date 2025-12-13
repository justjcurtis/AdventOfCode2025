package utils

func PrintResults(day int, results []string) {
	println("=------ Day " + Itoa(day) + " ------=")
	for i, result := range results {
		println("Part " + Itoa(i+1) + ": " + result)
	}
	println("=-------------------=")
}
