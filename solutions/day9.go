package solutions

import (
	. "AdventOfCode2025/models"
	"AdventOfCode2025/utils"
	"strings"
	"sync"
	"sync/atomic"
)

func parseDay9(input []string) []Vec {
	positions := make([]Vec, len(input))
	for i, line := range input {
		numStrs := strings.Split(line, ",")
		x := utils.Atoi(numStrs[0])
		y := utils.Atoi(numStrs[1])
		positions[i] = Vec{X: x, Y: y}
	}
	return positions
}

type Rect struct {
	left   int
	right  int
	top    int
	bottom int
}

func getRect(a, b Vec) Rect {
	left := a.X
	right := b.X
	top := a.Y
	bottom := b.Y

	if right < left {
		left, right = right, left
	}
	if bottom < top {
		top, bottom = bottom, top
	}

	return Rect{left, right, top, bottom}
}

func rectsOverlap(r0, r1 Rect) bool {
	return r0.left < r1.right &&
		r0.right > r1.left &&
		r0.top < r1.bottom &&
		r0.bottom > r1.top
}

func checkCollisions(positions []Vec, corner0, corner1 Vec, idx0, idx1 int) bool {
	r0 := getRect(corner0, corner1)
	for i := range positions {
		n := (i + 1) % len(positions)
		if i == idx0 || i == idx1 || n == idx0 || n == idx1 {
			continue
		}
		r1 := getRect(positions[i], positions[n])
		if rectsOverlap(r0, r1) {
			return false
		}
	}

	return true
}

func solveDay9(positions []Vec, part2 bool) int {
	var bestArea atomic.Int64
	upper := 1575000000
	lower := 1570000000
	if len(positions) < 100 {
		lower = 0
	}
	fn := func(i int) {
		a := positions[i]
		for j := i + 1; j < len(positions); j++ {
			b := positions[j]
			width := utils.IntAbs(a.X-b.X) + 1
			height := utils.IntAbs(a.Y-b.Y) + 1
			area := width * height
			if area <= int(bestArea.Load()) {
				continue
			}
			if part2 && (area < lower || area > upper || !checkCollisions(positions, a, b, i, j)) {
				continue
			}
			bestArea.Store(int64(area))
		}
	}
	utils.ParalleliseVoid(fn, len(positions))
	return int(bestArea.Load())
}

func Day9(input []string) []string {
	positions := parseDay9(input)
	var part2 int
	wg := sync.WaitGroup{}
	wg.Go(func() {
		part2 = solveDay9(positions, true)
	})
	part1 := solveDay9(positions, false)
	wg.Wait()
	return []string{utils.Itoa(part1), utils.Itoa(part2)}
}
