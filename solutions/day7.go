package solutions

import (
	"AdventOfCode2025/utils"
)

func solveDay7(input [][]byte) (int, int) {
	width := len(input[0])
	mid := width / 2
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
			t := timelines[j]
			timelines[j] = 0
			timelines[j-1] += t
			timelines[j+1] += t

			leftBlocked := line[j-1] == '|'
			rightBlocked := line[j+1] == '|'

			switch {
			case leftBlocked && rightBlocked:
				last := len(beams) - 1
				beams[b] = beams[last]
				beams = beams[:last]
				continue
			case rightBlocked:
				beams[b] = j - 1
				line[j-1] = '|'
				continue
			case leftBlocked:
				beams[b] = j + 1
				line[j+1] = '|'
				continue
			default:
				beams[b] = j + 1
				line[j+1] = '|'
				beams = append(beams, j-1)
				line[j-1] = '|'
			}
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
	return []string{utils.Itoa(part1), utils.Itoa(part2)}
}
