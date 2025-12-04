package solutions

import (
	"AdventOfCode2025/utils"
	"strconv"
)

func parseInputDay4(input []string) [][]rune {
	grid := make([][]rune, len(input))
	for i, line := range input {
		grid[i] = []rune(line)
	}
	return grid
}

func getNeighbors(x, y, width, height int) []struct{ dx, dy int } {
	neighbors := []struct{ dx, dy int }{}
	directions := []struct{ dx, dy int }{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
	}
	for _, d := range directions {
		nx, ny := x+d.dx, y+d.dy
		if nx >= 0 && nx < width && ny >= 0 && ny < height {
			neighbors = append(neighbors, d)
		}
	}
	return neighbors
}

func countOccupied(grid [][]rune, x, y int) int {
	count := 0
	neighbors := []struct{ dx, dy int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	for _, n := range neighbors {
		nx, ny := x+n.dx, y+n.dy
		if ny >= 0 && ny < len(grid) && nx >= 0 && nx < len(grid[ny]) {
			char := grid[ny][nx]
			if char == '@' {
				count++
			}
		}
	}
	return count
}

func solveDay4Part1(grid [][]rune) int {
	fn := func(i int) int {
		x, y := utils.OneDTwoD(i, len(grid[0]))
		if grid[y][x] == '@' {
			occupied := countOccupied(grid, x, y)
			if occupied < 4 {
				return 1
			}
		}
		return 0
	}
	return utils.Parallelise(utils.IntAcc, fn, len(grid)*len(grid[0]))
}

func solveDay4Part2(grid [][]rune) int {
	width := len(grid[0])
	height := len(grid)
	totalSize := height * width
	count := 0
	updates := make([]int, totalSize)
	updated := make([]int, totalSize)
	newGrid := make([][]rune, height)
	for y := 0; y < height; y++ {
		newGrid[y] = make([]rune, width)
		for x := 0; x < width; x++ {
			newGrid[y][x] = grid[y][x]
		}
	}

	for i := range updated {
		updates[i] = -1
	}
	checkUpdates := false
	fn := func(i int) int {
		var x, y int
		if checkUpdates {
			x, y = utils.OneDTwoD(updated[i], width)
		} else {
			x, y = utils.OneDTwoD(i, width)
		}
		char := grid[y][x]
		if char == '@' {
			occupied := countOccupied(grid, x, y)
			if occupied < 4 {
				newGrid[y][x] = '.'
				n := getNeighbors(x, y, width, height)
				for _, neighbor := range n {
					nx, ny := x+neighbor.dx, y+neighbor.dy
					if grid[ny][nx] == '@' {
						key := utils.TwoDToOneD(nx, ny, width)
						updates[key] = key
					}
				}
				return 1
			}
		}
		return 0
	}

	for {
		newCount := utils.Parallelise(utils.IntAcc, fn, len(updated))
		if newCount == 0 {
			return count
		}

		filtered := []int{}
		for _, idx := range updates {
			if idx != -1 {
				updates[idx] = -1
				filtered = append(filtered, idx)
			}
		}
		updated = filtered
		grid = newGrid

		checkUpdates = true
		count += newCount
	}
}

func Day4(input []string) []string {
	grid := parseInputDay4(input)
	part1 := solveDay4Part1(grid)
	part2 := solveDay4Part2(grid)
	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
