package solutions

import (
	"AdventOfCode2025/utils"
	"strconv"
	"strings"
	"sync"
)

func parseDay2(input []string) [][]int {
	rangeStrings := strings.Split(input[0], ",")
	ranges := make([][]int, len(rangeStrings))
	for i, r := range rangeStrings {
		bounds := strings.Split(r, "-")
		min, _ := strconv.Atoi(bounds[0])
		max, _ := strconv.Atoi(bounds[1])
		ranges[i] = []int{min, max}
	}
	return ranges
}

func solveDay2(ranges [][]int, unlocked bool) int {
	total := 0
	for _, r := range ranges {
		total += r[1] - r[0] + 1
	}
	toCheck := make([]int, total)
	added := 0
	for _, r := range ranges {
		a, b := r[0], r[1]
		for j := 0; j <= b-a; j++ {
			toCheck[added] = a + j
			added++
		}
	}
	fnSub := func(n string) bool {
		length := len(n)
		if !unlocked && length%2 != 0 {
			return false
		}
		mid := length / 2
		start := mid
		if unlocked {
			start = 1
		}
		for i := start; i <= mid; i++ {
			if length%i != 0 {
				continue
			}
			pattern := n[:i]
			match := true
			for j := i; j < length; j += i {
				if n[j:j+i] != pattern {
					match = false
					break
				}
			}
			if match {
				return match
			}
		}
		return false
	}
	fn := func(i int) int {
		n := toCheck[i]
		if fnSub(strconv.Itoa(n)) {
			return n
		}
		return 0
	}
	return utils.Parallelise(utils.IntAcc, fn, total)
}

func Day2(input []string) []string {
	parsed := parseDay2(input)
	wg := sync.WaitGroup{}
	wg.Add(2)
	var day2Part1, day2Part2 int
	go func() {
		defer wg.Done()
		day2Part1 = solveDay2(parsed, false)
	}()
	go func() {
		defer wg.Done()
		day2Part2 = solveDay2(parsed, true)
	}()
	wg.Wait()
	return []string{strconv.Itoa(day2Part1), strconv.Itoa(day2Part2)}
}
