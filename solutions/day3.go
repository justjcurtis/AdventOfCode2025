package solutions

import (
	"AdventOfCode2025/utils"
	"sync"
)

func parseDay3Line(line string) []int {
	result := make([]int, len(line))
	for i, char := range line {
		result[i] = int(char - '0')
	}
	return result
}

func parseDay3Input(input []string) [][]int {
	result := make([][]int, len(input))
	for i := range input {
		result[i] = parseDay3Line(input[i])
	}
	return result
}

func solveDay3(banks [][]int, allowed int) int {
	fn := func(i int) int {
		bank := banks[i]
		startIdx := 0
		buffer := make([]int, allowed)
		for current := range allowed {
			for j := startIdx; j < len(bank)-(allowed-current-1); j++ {
				num := bank[j]
				if num > buffer[current] {
					buffer[current] = num
					startIdx = j + 1
				}
			}
		}
		r := make([]rune, allowed)
		for j, val := range buffer {
			r[j] = rune(val + '0')
		}
		return utils.Atoi(string(r))
	}
	return utils.Parallelise(utils.IntAcc, fn, len(banks))
}

func Day3(input []string) []string {
	banks := parseDay3Input(input)
	wg := sync.WaitGroup{}
	var part1 int
	wg.Go(func() {
		part1 = solveDay3(banks, 2)
	})
	part2 := solveDay3(banks, 12)
	wg.Wait()
	return []string{utils.Itoa(part1), utils.Itoa(part2)}
}
