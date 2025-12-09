package solutions

import (
	. "AdventOfCode2025/models"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func parseDay8Input(input []string) []Vec3 {
	positions := make([]Vec3, len(input))
	for i, line := range input {
		numStrs := strings.Split(line, ",")
		nums := make([]int, len(numStrs))
		for j, str := range numStrs {
			nums[j], _ = strconv.Atoi(str)
		}
		positions[i] = Vec3{
			X: nums[0],
			Y: nums[1],
			Z: nums[2],
		}
	}

	return positions
}

type Pair struct {
	A        Vec3
	B        Vec3
	I        int
	J        int
	Distance float64
}

type Range struct {
	Min int
	Max int
}

type BBox struct {
	X Range
	Y Range
	Z Range
}

func getThreshold(positions []Vec3) float64 {
	n := len(positions)
	bb := BBox{}
	for i := range n {
		p := positions[i]
		if p.X < bb.X.Min {
			bb.X.Min = p.X
		}
		if p.X > bb.X.Max {
			bb.X.Max = p.X
		}
		if p.Y < bb.Y.Min {
			bb.Y.Min = p.Y
		}
		if p.Y > bb.Y.Max {
			bb.Y.Max = p.Y
		}
		if p.Z < bb.Z.Min {
			bb.Z.Min = p.Z
		}
		if p.Z > bb.Z.Max {
			bb.Z.Max = p.Z
		}
	}
	dx := float64(bb.X.Max - bb.X.Min)
	dy := float64(bb.Y.Max - bb.Y.Min)
	dz := float64(bb.Z.Max - bb.Z.Min)
	avg := (dx + dy + dz) / 3.0
	coef := 0.033
	if n < 1000 {
		coef = 0.5
	}
	return avg * avg * coef
}

func makePairs(positions []Vec3, threshold float64) ([]*Pair, []int) {
	n := len(positions)
	pairs := []*Pair{}
	index := 0
	for i := range n {
		a := positions[i]
		for j := i + 1; j < n; j++ {
			b := positions[j]
			d := a.FastDistance(b)
			if d > threshold {
				continue
			}
			pair := &Pair{
				A:        a,
				B:        b,
				I:        i,
				J:        j,
				Distance: d,
			}
			pairs = append(pairs, pair)
			index++
		}
	}

	order := make([]int, len(pairs))
	for i := range order {
		order[i] = i
	}
	sort.Slice(order, func(i, j int) bool {
		return pairs[order[i]].Distance < pairs[order[j]].Distance
	})

	return pairs, order
}

func solveDay8(pairs []*Pair, order []int, n int, part2 bool) int {
	stop := 1000
	if n < stop {
		stop = 10
	}
	circuits := make([]int, n)
	currentCircuit := 1
	connectionCounts := make(map[int]int)

	lastxset := [2]int{}
	for idx, o := range order {
		if idx == stop && !part2 {
			break
		}

		pair := pairs[o]
		i := pair.I
		j := pair.J
		lastxset[0] = pair.A.X
		lastxset[1] = pair.B.X
		ccCheck := circuits[i]
		if circuits[i] == 0 && circuits[j] == 0 {
			circuits[i] = currentCircuit
			circuits[j] = currentCircuit
			connectionCounts[currentCircuit] = 2
			ccCheck = currentCircuit
			currentCircuit++
		} else if circuits[i] != 0 && circuits[j] == 0 {
			circuits[j] = circuits[i]
			connectionCounts[circuits[i]]++
		} else if circuits[i] == 0 && circuits[j] != 0 {
			circuits[i] = circuits[j]
			connectionCounts[circuits[j]]++
			ccCheck = circuits[j]
		} else if circuits[i] != 0 && circuits[j] != 0 && circuits[i] != circuits[j] {
			oldCircuit := circuits[j]
			newCircuit := circuits[i]
			for k := range circuits {
				if circuits[k] == oldCircuit {
					circuits[k] = newCircuit
				}
			}
			connectionCounts[newCircuit] += connectionCounts[oldCircuit]
			delete(connectionCounts, oldCircuit)
		}
		if part2 && connectionCounts[ccCheck] == n {
			return lastxset[1] * lastxset[0]
		}
	}

	top3 := make([]int, 3)
	for _, count := range connectionCounts {
		if count > top3[0] {
			top3[2] = top3[1]
			top3[1] = top3[0]
			top3[0] = count
		} else if count > top3[1] {
			top3[2] = top3[1]
			top3[1] = count
		} else if count > top3[2] {
			top3[2] = count
		}
	}

	return top3[0] * top3[1] * top3[2]
}

func Day8(input []string) []string {
	positions := parseDay8Input(input)
	pairs, order := makePairs(positions, getThreshold(positions))
	var part1 int
	wg := sync.WaitGroup{}
	wg.Go(func() {
		part1 = solveDay8(pairs, order, len(positions), false)
	})
	part2 := solveDay8(pairs, order, len(positions), true)
	wg.Wait()
	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
