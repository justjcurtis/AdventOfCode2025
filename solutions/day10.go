package solutions

import (
	"AdventOfCode2025/utils"
	"math"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type Machine10 struct {
	Target   []bool
	Buttons  [][]int
	Joltages []int
}

func parseLine10(line string) Machine10 {
	sections := make([]string, 3)
	prev := 0
	for i, char := range line {
		if prev == 0 && char == '(' {
			sections[0] = line[prev+1 : i-2]
			prev = i
		}
		if prev > 0 && char == '{' {
			sections[1] = line[prev : i-1]
			sections[2] = line[i+1 : len(line)-1]
		}
	}
	target := make([]bool, len(sections[0]))
	for i, char := range sections[0] {
		if char == '#' {
			target[i] = true
		}
	}

	buttonStr := strings.Split(sections[1], " ")
	buttons := make([][]int, len(buttonStr))
	for i, bStr := range buttonStr {
		str := bStr[1 : len(bStr)-1]
		numStrs := strings.Split(str, ",")
		buttons[i] = make([]int, len(numStrs))
		for j, nStr := range numStrs {
			num, _ := strconv.Atoi(nStr)
			buttons[i][j] = num
		}
	}
	joltageStrs := strings.Split(sections[2], ",")
	joltages := make([]int, len(joltageStrs))
	for i, jStr := range joltageStrs {
		num, _ := strconv.Atoi(jStr)
		joltages[i] = num
	}

	return Machine10{
		Target:   target,
		Buttons:  buttons,
		Joltages: joltages,
	}
}

func parseDay10(input []string) []Machine10 {
	machines := make([]Machine10, len(input))
	fn := func(i int) {
		machines[i] = parseLine10(input[i])
	}
	utils.ParalleliseVoid(fn, len(input))
	return machines
}

func arrEquals[T comparable](a, joltages []T) bool {
	if len(a) != len(joltages) {
		return false
	}
	for i := range a {
		if a[i] != joltages[i] {
			return false
		}
	}
	return true
}

func generateCombosNoRepeat(length, opts int) [][]int {
	result := [][]int{}
	if length == 1 {
		for i := range opts {
			result = append(result, []int{i})
		}
		return result
	}
	subCombos := generateCombosNoRepeat(length-1, opts)
	for _, sub := range subCombos {
		used := map[int]bool{}
		for _, v := range sub {
			used[v] = true
		}
		for i := range opts {
			if !used[i] {
				newCombo := append([]int{}, sub...)
				newCombo = append(newCombo, i)
				result = append(result, newCombo)
			}
		}
	}
	return result
}

func solveDay10Part1(machines []Machine10) int {
	fn := func(i int) int {
		machine := machines[i]
		for presses := 1; presses <= len(machine.Buttons); presses++ {
			combos := generateCombosNoRepeat(presses, len(machine.Buttons))
			for _, combo := range combos {
				state := make([]bool, len(machine.Target))
				for _, buttonIndex := range combo {
					button := machine.Buttons[buttonIndex]
					for _, toggleIndex := range button {
						state[toggleIndex] = !state[toggleIndex]
					}
				}
				if arrEquals(state, machine.Target) {
					return presses
				}
			}
		}
		return -1
	}
	return utils.Parallelise(utils.IntAcc, fn, len(machines))
}

func swapRow(buttonMatrix [][]int, joltages []int, i, j int) {
	if i != j {
		buttonMatrix[i], buttonMatrix[j] = buttonMatrix[j], buttonMatrix[i]
		joltages[i], joltages[j] = joltages[j], joltages[i]
	}
}

func swapCol(buttonMatrix [][]int, bounds []int, i, j int) {
	if i != j {
		for k := range buttonMatrix {
			buttonMatrix[k][i], buttonMatrix[k][j] = buttonMatrix[k][j], buttonMatrix[k][i]
		}
		bounds[i], bounds[j] = bounds[j], bounds[i]
	}
}

func reduceRow(buttonMatrix [][]int, joltages []int, i, j int) {
	x := buttonMatrix[i][i]
	y := -buttonMatrix[j][i]
	d := utils.GCD(x, y)
	for k := range buttonMatrix[i] {
		buttonMatrix[j][k] = (y*buttonMatrix[i][k] + x*buttonMatrix[j][k]) / d
	}
	joltages[j] = (y*joltages[i] + x*joltages[j]) / d
}

func reduceButtonMatrix(buttonMatrix [][]int, joltages, bounds []int) ([][]int, []int, []int) {
	nRows := len(buttonMatrix)
	if nRows == 0 {
		return buttonMatrix, joltages, bounds
	}
	nCols := len(buttonMatrix[0])

	for i := range nCols {
		buf := []int{}
		k := i
		for len(buf) == 0 && k < nCols {
			swapCol(buttonMatrix, bounds, i, k)
			buf = []int{}
			for j := i; j < nRows; j++ {
				if buttonMatrix[j][i] != 0 {
					buf = append(buf, j)
				}
			}
			k++
		}
		if len(buf) == 0 {
			break
		}
		swapRow(buttonMatrix, joltages, i, buf[0])
		for j := i + 1; j < nRows; j++ {
			reduceRow(buttonMatrix, joltages, i, j)
		}
	}

	buf := []int{}
	for i := range buttonMatrix {
		nonZero := false
		for _, v := range buttonMatrix[i] {
			if v != 0 {
				nonZero = true
				break
			}
		}
		if nonZero {
			buf = append(buf, i)
		}
	}
	newButtonMatrix := [][]int{}
	newJoltages := []int{}
	for _, i := range buf {
		newButtonMatrix = append(newButtonMatrix, buttonMatrix[i])
		newJoltages = append(newJoltages, joltages[i])
	}

	for i := len(newButtonMatrix) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			reduceRow(newButtonMatrix, newJoltages, i, j)
		}
	}

	return newButtonMatrix, newJoltages, bounds
}

func minimizeButtonPressSum(buttonMatrix [][]int, joltages, bounds []int) int {
	r := len(buttonMatrix)
	n := len(buttonMatrix[0])
	free := utils.IntMax(n-r, 0)

	freeCombos := generateFreeParamCombos(free, bounds)

	minSum := math.MaxInt32

	for _, fc := range freeCombos {
		valid := true
		depVars := make([]int, r)
		for i := range r {
			sum := 0
			for j := range free {
				sum += fc[j] * buttonMatrix[i][r+j]
			}
			diff := joltages[i] - sum
			if diff%buttonMatrix[i][i] != 0 {
				valid = false
				break
			}
			dep := diff / buttonMatrix[i][i]
			if dep < 0 {
				valid = false
				break
			}
			depVars[i] = dep
		}

		if !valid {
			continue
		}

		total := 0
		for _, v := range depVars {
			total += v
		}
		for _, v := range fc {
			total += v
		}

		if total < minSum {
			minSum = total
		}
	}

	if minSum == math.MaxInt32 {
		return -1
	}
	return minSum
}

func generateFreeParamCombos(nparam int, bounds []int) [][]int {
	if nparam == 0 {
		return [][]int{{}}
	}
	ret := [][]int{}
	bound := bounds[len(bounds)-nparam]
	for i := 0; i <= bound; i++ {
		for _, sub := range generateFreeParamCombos(nparam-1, bounds) {
			combo := append([]int{i}, sub...)
			ret = append(ret, combo)
		}
	}
	return ret
}

func buildMatrix(machine Machine10) ([][]int, []int) {
	bm := make([][]int, len(machine.Joltages))
	for i := range bm {
		bm[i] = make([]int, len(machine.Buttons))
		for j := range machine.Buttons {
			if slices.Contains(machine.Buttons[j], i) {
				bm[i][j] = 1
			}
		}
	}

	bounds := make([]int, len(machine.Buttons))
	for i := range machine.Buttons {
		minPresses := -1
		for _, toggleIndex := range machine.Buttons[i] {
			if minPresses == -1 || machine.Joltages[toggleIndex] < minPresses {
				minPresses = machine.Joltages[toggleIndex]
			}
		}
		bounds[i] = minPresses
	}

	return bm, bounds
}

func solveDay10Part2(machines []Machine10) int {
	fn := func(i int) int {
		machine := machines[i]
		buttonMatrix, bounds := buildMatrix(machine)
		reducedMatrix, reducedTarget, freeParamBounds := reduceButtonMatrix(buttonMatrix, machine.Joltages, bounds)
		return minimizeButtonPressSum(reducedMatrix, reducedTarget, freeParamBounds)
	}
	return utils.Parallelise(utils.IntAcc, fn, len(machines))
}

func Day10(input []string) []string {
	machines := parseDay10(input)
	var part1 int
	wg := sync.WaitGroup{}
	wg.Go(func() {
		part1 = solveDay10Part1(machines)
	})
	part2 := solveDay10Part2(machines)
	wg.Wait()
	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
