package solutions

import (
	"AdventOfCode2025/utils"
	"sync"
)

func parseInputDay4(input []string) [][]rune {
	fn := func(i int) []rune {
		return []rune(input[i])
	}
	return utils.ParalleliseMap(fn, len(input))
}

var dirs = []struct{ dx, dy int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

func solveDay4(grid [][]rune, runForever bool) int {
	width := len(grid[0])
	height := len(grid)
	totalSize := height * width
	count := 0
	g := make([]byte, totalSize)
	stack := make([]int, totalSize)

	pos := 0
	for y := range height {
		for x := range width {
			char := grid[y][x]
			g[pos] = byte(char)
			stack[pos] = pos
			pos++
		}
	}

	for len(stack) > 0 {
		i := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		x, y := utils.OneDTwoD(i, width)
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
	wg := sync.WaitGroup{}
	var part1 int
	wg.Go(func() {
		part1 = solveDay4(grid, false)
	})
	part2 := solveDay4(grid, true)
	wg.Wait()
	return []string{utils.Itoa(part1), utils.Itoa(part2)}
}
