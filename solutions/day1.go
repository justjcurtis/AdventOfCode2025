// Package solutions contains solutions to the coding challenges.
package solutions

import (
	"AdventOfCode2025/utils"
	"sync"
)

func parseDay1Line(line string) int {
	value := utils.Atoi(line[1:])
	if line[0] == 'L' {
		value = -value
	}
	return value
}

func parseDay1(input []string) []int {
	fn := func(i int) int {
		return parseDay1Line(input[i])
	}
	return utils.ParalleliseMap(fn, len(input))
}

func solveDay1Part1(moves []int) int {
	value := 50
	count := 0
	for _, move := range moves {
		value += move
		if value%100 == 0 {
			count++
		}
	}
	return count
}

func solveDay1Part2(moves []int) int {
	value := 50
	count := 0
	for _, move := range moves {
		prev := value
		absMove := utils.IntAbs(move)
		count += absMove / 100
		rem := move % 100
		after := prev + rem
		value = after % 100

		if absMove > 100 {
			if rem == 0 {
				if value == 0 {
					count--
				}
			}
		}

		if value != 0 && prev != 0 {
			if utils.GetSign(prev) != utils.GetSign(after) {
				count++
			} else if after/100 != prev/100 {
				count++
			}
		}
		if value == 0 {
			count++
		}
	}
	return count
}

func Day1(input []string) []string {
	parsed := parseDay1(input)
	wg := sync.WaitGroup{}
	var part1 int
	wg.Go(func() {
		part1 = solveDay1Part1(parsed)
	})
	part2 := solveDay1Part2(parsed)
	wg.Wait()
	return []string{utils.Itoa(part1), utils.Itoa(part2)}
}
