package solutions

import (
	"AdventOfCode2025/utils"
	"strings"
)

type Space struct {
	area     int
	required []int
}

func parseInputDay12(input []string) ([]int, []Space) {
	blockIndex := -1
	blocks := []int{}
	spaces := []Space{}
	for _, line := range input {
		_ = line
		if strings.Contains(line, "#") {
			blocks[blockIndex] += strings.Count(line, "#")
			continue
		}
		if strings.Contains(line, "x") {
			strs := strings.Split(line, ": ")
			dimensions := strings.Split(strs[0], "x")
			width := utils.Atoi(dimensions[0])
			height := utils.Atoi(dimensions[1])
			space := Space{area: width * height, required: []int{}}
			countStrs := strings.SplitSeq(strs[1], " ")
			for countStr := range countStrs {
				space.required = append(space.required, utils.Atoi(countStr))
			}
			spaces = append(spaces, space)
			continue
		}
		if strings.Contains(line, ":") {
			blocks = append(blocks, 0)
			blockIndex++
		}
	}
	return blocks, spaces
}

func solveDay12(blocks []int, spaces []Space) int {
	fn := func(i int) int {
		space := spaces[i]
		sizeRequired := 0
		for j, req := range space.required {
			sizeRequired += req * blocks[j]
		}
		if sizeRequired*115 < space.area*100 {
			return 1
		}
		return 0
	}
	return utils.Parallelise(utils.IntAcc, fn, len(spaces))
}

func Day12(input []string) []string {
	blocks, spaces := parseInputDay12(input)
	part1 := solveDay12(blocks, spaces)
	return []string{utils.Itoa(part1)}
}
