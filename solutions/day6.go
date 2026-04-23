package solutions

import (
	"AdventOfCode2025/utils"
	"fmt"
	"sync"
)

type WorkbookColumn struct {
	nums [][]byte
	op   byte
}

func newColBuffer(height int) [][]byte {
	nums := make([][]byte, height)
	for i := range height {
		nums[i] = []byte{}
	}
	return nums
}

func getCols(input []string) [][2]int {
	width := len(input[0])
	i := len(input) - 1
	result := [][2]int{}
	lastGap := 0
	for j := 1; j < width; j++ {
		char := input[i][j]
		if char == ' ' {
			continue
		}
		result = append(result, [2]int{lastGap, (j - 1) - lastGap})
		lastGap = j
	}
	result = append(result, [2]int{lastGap, width - lastGap})
	return result
}

func parseDay6Col(input []string, start, width, height int) WorkbookColumn {
	colBuffer := newColBuffer(height - 1)
	op := byte(' ')
	for j := start; j < start+width; j++ {
		vert := make([]byte, height-1)
		for i := 0; i < height-1; i++ {
			char := input[i][j]
			vert[i] = char
		}
		for v, char := range vert {
			colBuffer[v] = append(colBuffer[v], char)
		}
		opChar := input[height-1][j]
		if opChar != ' ' {
			op = opChar
		}
	}
	return WorkbookColumn{nums: colBuffer, op: op}

}

func parseDay6Input(input []string) []WorkbookColumn {
	height := len(input)
	colsInfo := getCols(input)
	fn := func(i int) WorkbookColumn {
		colInfo := colsInfo[i]
		start := colInfo[0]
		colWidth := colInfo[1]
		return parseDay6Col(input, start, colWidth, height)
	}
	return utils.ParalleliseMap(fn, len(colsInfo))
}

func doOperation(a int, b int, op byte) int {
	switch op {
	case '+':
		return a + b
	case '*':
		return a * b
	default:
		panic(fmt.Sprintf("doOperation: unknown operator %c", op))
	}
}

func getNums(numsRaw [][]byte, cephalopod bool) []int {
	if !cephalopod {
		nums := make([]int, len(numsRaw))
		for i, s := range numsRaw {
			numBuffer := 0
			for _, char := range s {
				if char == ' ' {
					continue
				}
				digit := int(char - '0')
				numBuffer = numBuffer*10 + digit
			}
			nums[i] = numBuffer
		}
		return nums
	}
	width := len(numsRaw[0])
	nums := make([]int, width)
	for j := range width {
		numBuffer := 0
		for i := range numsRaw {
			char := numsRaw[i][j]
			if char == ' ' {
				continue
			}
			digit := int(char - '0')
			numBuffer = numBuffer*10 + digit
		}
		nums[j] = numBuffer
	}
	return nums
}

func solveDay6(columns []WorkbookColumn, cephalopod bool) int {
	fn := func(i int) int {
		col := columns[i]
		nums := getNums(col.nums, cephalopod)
		result := nums[0]
		for j := 1; j < len(nums); j++ {
			result = doOperation(result, nums[j], col.op)
		}
		return result
	}
	return utils.Parallelise(utils.IntAcc, fn, len(columns))
}

func Day6(input []string) []string {
	parsed := parseDay6Input(input)
	var part1 int
	wg := sync.WaitGroup{}
	wg.Go(func() {
		part1 = solveDay6(parsed, false)
	})
	part2 := solveDay6(parsed, true)
	wg.Wait()
	return []string{utils.Itoa(part1), utils.Itoa(part2)}
}
