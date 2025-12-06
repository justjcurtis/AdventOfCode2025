package solutions

import (
	"AdventOfCode2025/utils"
	"fmt"
	"strconv"
)

type WorkbookColumn struct {
	nums [][]byte
	op   byte
}

var colBufferCache [][]byte

func newColBuffer(height int) [][]byte {
	if colBufferCache == nil {
		nums := make([][]byte, height)
		for i := range height {
			nums[i] = []byte{}
		}
		colBufferCache = nums
		return nums
	}
	for i := range height {
		colBufferCache[i] = colBufferCache[i][:0]
	}
	return colBufferCache
}

func parseDay6Input(input []string) []WorkbookColumn {
	columns := []WorkbookColumn{}
	width := len(input[0])
	height := len(input)
	colBuffer := newColBuffer(height - 1)
	op := byte(' ')
	for j := range width {
		isGap := true
		vert := make([]byte, height-1)
		for i := 0; i < height-1; i++ {
			char := input[i][j]
			vert[i] = char
			if char != ' ' {
				isGap = false
			}
		}
		if isGap {
			columns = append(columns, WorkbookColumn{nums: utils.Clone2D(colBuffer), op: op})
			colBuffer = newColBuffer(height - 1)
			continue
		}
		for v, char := range vert {
			colBuffer[v] = append(colBuffer[v], char)
		}
		opChar := input[height-1][j]
		if opChar != ' ' {
			op = opChar
		}
	}
	columns = append(columns, WorkbookColumn{nums: utils.Clone2D(colBuffer), op: op})
	return columns
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

func getNums(numStrs [][]byte, cephalopod bool) []int {
	if !cephalopod {
		nums := make([]int, len(numStrs))
		for i, s := range numStrs {
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
	width := len(numStrs[0])
	nums := make([]int, width)
	for j := range width {
		numBuffer := 0
		for i := range numStrs {
			char := numStrs[i][j]
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
	part1 := solveDay6(parsed, false)
	part2 := solveDay6(parsed, true)
	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
