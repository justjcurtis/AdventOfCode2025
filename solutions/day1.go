// Package solutions contains solutions to the coding challenges.
package solutions

import (
	"AdventOfCode2025/utils"
	"math"
	"strconv"
	"sync"
)

func parseDay1Line(line string) int {
	factor := 1
	if line[0] == 'L' {
		factor = -1
	}
	value, _ := strconv.Atoi(line[1:])
	return factor * value
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

func getSign(x int) int {
	if x < 0 {
		return -1
	}
	return 1
}

func solveDay1Part2(moves []int) int {
	value := 50
	count := 0
	for _, move := range moves {
		prev := value
		absMove := int(math.Abs(float64(move)))
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
			if getSign(prev) != getSign(after) {
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
	wg.Add(2)
	var part1, part2 int
	go func() {
		defer wg.Done()
		part1 = solveDay1Part1(parsed)
	}()
	go func() {
		defer wg.Done()
		part2 = solveDay1Part2(parsed)
	}()
	wg.Wait()
	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
