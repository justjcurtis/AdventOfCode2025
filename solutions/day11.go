package solutions

import (
	"AdventOfCode2025/utils"
	"strconv"
	"strings"
	"sync"
)

func parseDay11(input []string) map[string][]string {
	devices := make(map[string][]string)
	for _, line := range input {
		parts := strings.Split(line, ": ")
		deviceName := parts[0]
		outputs := strings.Split(parts[1], " ")
		devices[deviceName] = outputs
	}

	return devices
}

func solveDay11Part1(devices map[string][]string, start, end string) int {
	solutionCache := sync.Map{}
	var recurse func(current string, visited map[string]bool) int
	fn := func(i int) int {
		device := devices[start][i]
		visited := make(map[string]bool)
		visited[start] = true
		return recurse(device, visited)
	}
	recurse = func(current string, visited map[string]bool) int {
		if val, ok := solutionCache.Load(current); ok {
			return val.(int)
		}
		if current == end {
			return 1
		}
		total := 0
		for _, next := range devices[current] {
			if visited[next] {
				continue
			}
			visited[next] = true
			total += recurse(next, visited)
			visited[next] = false
		}
		solutionCache.Store(current, total)
		return total
	}
	return utils.Parallelise(utils.IntAcc, fn, len(devices[start]))
}

func solveDay11Part2(devices map[string][]string) int {
	parts := [][]string{
		{"svr", "fft"},
		{"fft", "dac"},
		{"dac", "out"},
	}
	fn := func(i int) int {
		start := parts[i][0]
		end := parts[i][1]
		return solveDay11Part1(devices, start, end)
	}
	return utils.Parallelise(utils.MultAcc, fn, len(parts))
}

func Day11(input []string) []string {
	devices := parseDay11(input)
	var part1 int
	wg := sync.WaitGroup{}
	wg.Go(func() {
		part1 = solveDay11Part1(devices, "you", "out")
	})
	part2 := solveDay11Part2(devices)
	wg.Wait()
	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
