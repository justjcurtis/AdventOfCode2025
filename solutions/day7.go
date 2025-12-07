package solutions

import (
	"AdventOfCode2025/utils"
	"strconv"
)

func solveDay7(input [][]byte) (int, int) {
	width := len(input[0])
	mid := (width / 2)
	beams := []int{mid}
	splitCount := 0
	timelines := make([]int, width)
	timelines[mid] = 1
	for i := 2; i < len(input); i++ {
		line := input[i]
		for b := len(beams) - 1; b >= 0; b-- {
			j := beams[b]
			if j <= 0 || j >= width-1 {
				continue
			}
			if line[j] != '^' {
				if i < len(input)-1 {
					if input[i+1][j] != '^' {
						input[i+1][j] = '|'
					}
				}
				continue
			}
			splitCount++
			timelines[j-1] += timelines[j]
			timelines[j+1] += timelines[j]
			timelines[j] = 0
			if line[j+1] == '|' {
				if line[j-1] == '|' {
					beams = append(beams[:b], beams[b+1:]...)
					continue
				}
				beams[b] = j - 1
				line[j-1] = '|'
				continue
			}
			beams[b] = j + 1
			line[j+1] = '|'
			if line[j-1] == '|' {
				continue
			}
			beams = append(beams, j-1)
			line[j-1] = '|'
		}
	}
	timelineCount := 0
	for _, v := range timelines {
		timelineCount += v
	}
	return splitCount, timelineCount
}

func Day7(input []string) []string {
	part1, part2 := solveDay7(utils.ConvertToByteMatrix(input))
	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
