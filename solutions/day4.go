package solutions

import (
	"AdventOfCode2025/utils"
	"strconv"
	"sync"
)

func parseInputDay4(input []string) [][]rune {
	grid := make([][]rune, len(input))
	for i, line := range input {
		grid[i] = []rune(line)
	}
	return grid
}

var dirs = []struct{ dx, dy int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

func solveDay4(grid [][]rune, runForever bool) int {
	width := len(grid[0])
	height := len(grid)
	totalSize := height * width
	count := 0
	g := make([]byte, totalSize)
	stack := make([]int, totalSize)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			char := grid[y][x]
			key := utils.TwoDToOneD(x, y, width)
			g[key] = byte(char)
			stack[key] = key
		}
	}

	for len(stack) > 0 {
		i := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		var x, y int
		x, y = utils.OneDTwoD(i, width)
		char := g[i]
		if char == '@' {

			occupied := 0
			for _, d := range dirs {
				nx, ny := x+d.dx, y+d.dy
				if nx >= 0 && nx < width && ny >= 0 && ny < height {
					nkey := utils.TwoDToOneD(nx, ny, width)
					if g[nkey] == '@' {
						occupied++
					}
				}
			}
			if occupied < 4 {
				count++
				if !runForever {
					continue
				}
				g[i] = '.'
				for _, d := range dirs {
					nx, ny := x+d.dx, y+d.dy
					if nx >= 0 && nx < width && ny >= 0 && ny < height {
						nkey := utils.TwoDToOneD(nx, ny, width)
						stack = append(stack, nkey)
					}
				}
				continue
			}
		}
	}

	return count
}

func Day4(input []string) []string {
	grid := parseInputDay4(input)
	_ = grid
	wg := sync.WaitGroup{}
	wg.Add(2)
	var part1, part2 int
	go func() {
		defer wg.Done()
		part1 = solveDay4(grid, false)
	}()
	go func() {
		defer wg.Done()
		part2 = solveDay4(grid, true)
	}()
	wg.Wait()
	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
