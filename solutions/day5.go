package solutions

import (
	"AdventOfCode2025/utils"
	"slices"
	"strings"
	"sync"
)

func condenseRanges(ranges [][2]int) [][2]int {
	slices.SortFunc(ranges, func(a, b [2]int) int {
		return a[0] - b[0]
	})

	for i := 0; i < len(ranges)-1; {
		current := ranges[i]
		next := ranges[i+1]

		if current[1] >= next[0]-1 {
			merged := [2]int{current[0], utils.IntMax(current[1], next[1])}
			ranges[i] = merged
			ranges = append(ranges[:i+1], ranges[i+2:]...)
		} else {
			i++
		}
	}

	return ranges
}

func parseInputDay5(input []string) ([][2]int, []int) {
	rangeGetting := true
	ranges := [][2]int{}
	ids := []int{}
	for _, line := range input {
		if line == "" {
			rangeGetting = false
			continue
		}
		if rangeGetting {
			parts := strings.Split(line, "-")
			a := utils.Atoi(parts[0])
			b := utils.Atoi(parts[1])
			ranges = append(ranges, [2]int{a, b})
			continue
		}
		id := utils.Atoi(line)
		ids = append(ids, id)
	}

	return condenseRanges(ranges), ids
}

func solveDay5Part1(ranges [][2]int, ids []int) int {
	fn := func(i int) int {
		id := ids[i]
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				return 1
			}
		}
		return 0
	}
	return utils.Parallelise(utils.IntAcc, fn, len(ids))
}

func solveDay5Part2(ranges [][2]int) int {
	countRange := func(i int) int {
		r := ranges[i]
		return r[1] - r[0] + 1
	}
	return utils.Parallelise(utils.IntAcc, countRange, len(ranges))
}

func Day5(input []string) []string {
	ranges, ids := parseInputDay5(input)
	var part1 int
	wg := sync.WaitGroup{}
	wg.Go(func() {
		part1 = solveDay5Part1(ranges, ids)
	})
	part2 := solveDay5Part2(ranges)
	wg.Wait()
	return []string{utils.Itoa(part1), utils.Itoa(part2)}
}
