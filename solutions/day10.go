package solutions

import (
	"AdventOfCode2025/utils"
	"bytes"
	"math"
	"math/bits"
	"slices"
	"sync"
)

const epsilon = 1e-9

func atoiByte(buf []byte) int {
	result := 0
	for _, c := range buf {
		result = result*10 + int(c-'0')
	}
	return result
}

func stripBrackets(input []byte) []byte {
	return input[1 : len(input)-1]
}

func parseInputDay10(input []string) ([][]uint16, []uint16, [][]int) {
	var switches [][]uint16
	var lights []uint16
	var jolts [][]int

	for _, line := range input {
		lineBytes := []byte(line)
		fields := bytes.Split(lineBytes, []byte(" "))

		lightFieldRaw := fields[0]
		lightPatternContent := stripBrackets(lightFieldRaw)

		var light uint16
		for i, char := range lightPatternContent {
			if char == '#' {
				light |= 1 << i
			}
		}
		lights = append(lights, light)

		switchFieldsRaw := fields[1 : len(fields)-1]
		currentSwitchRow := make([]uint16, len(switchFieldsRaw))

		for idx, field := range switchFieldsRaw {
			fieldContent := stripBrackets(field)

			for part := range bytes.SplitSeq(fieldContent, []byte(",")) {
				position := atoiByte(part)
				currentSwitchRow[idx] |= 1 << uint16(position)
			}
		}
		switches = append(switches, currentSwitchRow)

		joltFieldRaw := fields[len(fields)-1]
		joltValuesContent := stripBrackets(joltFieldRaw)

		joltSets := bytes.Split(joltValuesContent, []byte(","))
		currentJolts := make([]int, len(joltSets))

		for i, valuePart := range joltSets {
			currentJolts[i] = atoiByte(valuePart)
		}
		jolts = append(jolts, currentJolts)
	}

	return switches, lights, jolts
}

func solveDay10Part1(switches [][]uint16, lights []uint16) int {
	return utils.Parallelise(utils.IntAcc, func(lightIndex int) int {
		row := switches[lightIndex]

		maxBitLength := 0
		for _, state := range row {
			stateBits := bits.Len16(state)
			if stateBits > maxBitLength {
				maxBitLength = stateBits
			}
		}

		numStates := 1 << maxBitLength
		distanceMap := make([]int, numStates)
		for i := range distanceMap {
			distanceMap[i] = -1
		}
		distanceMap[0] = 0

		queue := []uint16{0}

		for idx := 0; idx < len(queue); idx++ {
			currentState := queue[idx]

			for _, switchConfig := range row {
				nextState := currentState ^ switchConfig

				if distanceMap[nextState] != -1 {
					continue
				}

				distanceMap[nextState] = distanceMap[currentState] + 1
				queue = append(queue, nextState)
			}
		}

		return distanceMap[lights[lightIndex]]
	}, len(lights))
}

func lpSimplexSolver(constraints [][]float64, objective []float64) (float64, []float64) {
	numConstraints := len(constraints)
	numVars := len(constraints[0]) - 1

	nonBasic := make([]int, numVars+1)
	for i := range numVars {
		nonBasic[i] = i
	}
	nonBasic[numVars] = -1

	basic := make([]int, numConstraints)
	for i := range numConstraints {
		basic[i] = numVars + i
	}

	tableau := make([][]float64, numConstraints+2)

	for i := range numConstraints {
		tableau[i] = make([]float64, numVars+2)
		tableau[i][numVars+1] = -1
		copy(tableau[i], constraints[i])
	}

	tableau[numConstraints] = make([]float64, numVars+2)
	copy(tableau[numConstraints], objective)
	tableau[numConstraints][numVars] = 0
	tableau[numConstraints][numVars+1] = 0

	tableau[numConstraints+1] = make([]float64, numVars+2)

	for i := range numConstraints {
		tableau[i][numVars], tableau[i][numVars+1] =
			tableau[i][numVars+1], tableau[i][numVars]
	}
	tableau[numConstraints+1][numVars] = 1

	pivot := func(row, col int) {
		scale := 1.0 / tableau[row][col]

		for i := range numConstraints + 2 {
			if i == row {
				continue
			}
			for j := range numVars + 2 {
				if j != col {
					tableau[i][j] -= tableau[row][j] * tableau[i][col] * scale
				}
			}
		}

		for j := range numVars + 2 {
			tableau[row][j] *= scale
		}

		for i := range numConstraints + 2 {
			tableau[i][col] *= -scale
		}

		tableau[row][col] = scale

		basic[row], nonBasic[col] = nonBasic[col], basic[row]
	}

	findPivotAndOptimize := func(phase int) bool {
		for {
			enterCol := -1
			minValue := math.Inf(1)

			for i := range numVars + 1 {
				if phase != 0 || nonBasic[i] != -1 {
					val := tableau[numConstraints+phase][i]
					if val < minValue || (val == minValue && nonBasic[i] < nonBasic[enterCol]) {
						enterCol = i
						minValue = val
					}
				}
			}

			if -epsilon < tableau[numConstraints+phase][enterCol] {
				return true
			}

			leaveRow := -1
			minRatio := math.Inf(1)

			for i := range numConstraints {
				if tableau[i][enterCol] > epsilon {
					ratio := tableau[i][numVars+1] / tableau[i][enterCol]
					if ratio < minRatio || (ratio == minRatio && basic[i] < basic[leaveRow]) {
						leaveRow = i
						minRatio = ratio
					}
				}
			}

			if leaveRow == -1 {
				return false
			}

			pivot(leaveRow, enterCol)
		}
	}

	startRow := 0
	minRHS := tableau[0][numVars+1]

	for i := 1; i < numConstraints; i++ {
		if tableau[i][numVars+1] < minRHS {
			startRow = i
			minRHS = tableau[i][numVars+1]
		}
	}

	if -epsilon > tableau[startRow][numVars+1] {
		pivot(startRow, numVars)

		if !findPivotAndOptimize(1) || tableau[numConstraints+1][numVars+1] < -epsilon {
			return math.Inf(-1), nil
		}
	}

	for i := range numConstraints {
		if basic[i] == -1 {
			enterCol := 0
			minVal := tableau[i][0]

			for j := 1; j < numVars; j++ {
				if tableau[i][j] < minVal || (tableau[i][j] == minVal && nonBasic[j] < nonBasic[enterCol]) {
					enterCol = j
					minVal = tableau[i][j]
				}
			}

			pivot(i, enterCol)
		}
	}

	if findPivotAndOptimize(0) {
		solution := make([]float64, numVars)
		for i := range numConstraints {
			if basic[i] >= 0 && basic[i] < numVars {
				solution[basic[i]] = tableau[i][numVars+1]
			}
		}

		result := 0.0
		for i := range numVars {
			result += objective[i] * solution[i]
		}

		return result, solution
	}

	return math.Inf(-1), nil
}

func branchAndBound(constraints [][]float64, objective []float64) int {
	bestValue := math.Inf(1)

	numVars := len(constraints[0]) - 1

	var branch func(currentConstraints [][]float64)

	branch = func(currentConstraints [][]float64) {
		val, solution := lpSimplexSolver(currentConstraints, objective)

		if bestValue < val+epsilon || math.IsInf(val, -1) {
			return
		}

		branchVar := -1
		branchVal := 0

		for i, sol := range solution {
			if epsilon < math.Abs(sol-math.Round(sol)) {
				branchVar = i
				branchVal = int(sol)
				break
			}
		}

		if branchVar == -1 {
			if val+epsilon < bestValue {
				bestValue = val
			}
		} else {
			newConstraint := make([]float64, numVars+1)
			newConstraint[numVars] = float64(branchVal)
			newConstraint[branchVar] = 1
			branch(append(currentConstraints, newConstraint))

			newConstraint = make([]float64, numVars+1)
			newConstraint[numVars] = float64(^branchVal)
			newConstraint[branchVar] = -1
			branch(append(currentConstraints, newConstraint))
		}
	}

	branch(constraints)

	return int(math.Round(bestValue))
}

func solveDay10Part2(switchBitmasks [][]uint16, joltLimits [][]int) int {
	processInstance := func(idx int) int {
		switchRow := switchBitmasks[idx]
		joltRow := joltLimits[idx]

		numSwitches := len(switchRow)
		maxBits := 0
		for _, mask := range switchRow {
			maxBits = max(maxBits, bits.Len16(mask))
		}

		constraintMatrix := make([][]float64, 2*maxBits+numSwitches)
		for r := range constraintMatrix {
			constraintMatrix[r] = make([]float64, numSwitches+1)
		}

		for col, mask := range switchRow {
			row := (2*maxBits + len(switchRow)) - 1 - col
			constraintMatrix[row][col] = -1

			for bit := 0; bit < maxBits; bit++ {
				if mask&(1<<bit) != 0 {
					constraintMatrix[bit][col] = 1
					constraintMatrix[bit+maxBits][col] = -1
				}
			}
		}

		for bit := range maxBits {
			constraintMatrix[bit][numSwitches] = float64(joltRow[bit])
			constraintMatrix[bit+maxBits][numSwitches] = -float64(joltRow[bit])
		}

		return branchAndBound(constraintMatrix, slices.Repeat([]float64{1}, numSwitches))
	}

	return utils.Parallelise(utils.IntAcc, processInstance, len(switchBitmasks))
}

func Day10(input []string) []string {
	switches, lights, jolts := parseInputDay10(input)
	var part1 int
	wg := sync.WaitGroup{}
	wg.Go(func() {
		part1 = solveDay10Part1(switches, lights)
	})
	part2 := solveDay10Part2(switches, jolts)
	wg.Wait()
	return []string{utils.Itoa(part1), utils.Itoa(part2)}
}
