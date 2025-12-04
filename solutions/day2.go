package solutions

import (
	"AdventOfCode2025/utils"
	"strconv"
	"strings"
	"sync"
)

func parseDay2(input []string) [][]int {
	rangeStrings := strings.Split(input[0], ",")
	ranges := make([][]int, len(rangeStrings))
	for i, r := range rangeStrings {
		bounds := strings.Split(r, "-")
		intMin, _ := strconv.Atoi(bounds[0])
		intMax, _ := strconv.Atoi(bounds[1])
		ranges[i] = []int{intMin, intMax}
	}
	return ranges
}

func lengthOfInt(n int) int {
	length := 0
	for n > 0 {
		n /= 10
		length++
	}
	return length
}

func intPow10(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}

func solveDay2Fast(ranges [][]int, unlocked bool) int {
	pow10 := []int{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000, 10000000000}

	intPow10 := func(n int) int { return pow10[n] }
	intMax := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	intMin := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}

	mobius := func(limit int) []int {
		mu := make([]int, limit+1)
		for i := range mu {
			mu[i] = 1
		}
		for i := 2; i <= limit; i++ {
			if mu[i] == 1 {
				for j := i; j <= limit; j += i {
					mu[j] *= -1
				}
				for j := i * i; j <= limit; j += i * i {
					mu[j] = 0
				}
			}
		}
		return mu
	}

	fn := func(n int) int {
		a, b := ranges[n][0], ranges[n][1]
		minLength := lengthOfInt(a)
		maxLength := lengthOfInt(b)
		var result int64 = 0

		for length := minLength; length <= maxLength; length++ {
			if !unlocked && length%2 != 0 {
				continue
			}

			sumByPeriod := make([]int64, length/2+1)
			maxPeriod := length / 2
			for p := 1; p <= maxPeriod; p++ {
				if length%p != 0 {
					continue
				}
				k := length / p
				if k < 2 {
					continue
				}

				var multiplier int64 = 0
				var factor int64 = 1
				for i := 0; i < k; i++ {
					multiplier += factor
					factor *= int64(intPow10(p))
				}

				low := intMax(intPow10(p-1), int((int64(a)+multiplier-1)/multiplier))
				high := intMin(b/int(multiplier), intPow10(p)-1)
				if low > high {
					continue
				}

				count := int64(high - low + 1)
				sum := count * int64(low+high) / 2
				sumByPeriod[p] = sum * multiplier
			}

			if unlocked {
				mu := mobius(maxPeriod)
				var primitiveSum int64 = 0
				for p := 1; p <= maxPeriod; p++ {
					if sumByPeriod[p] == 0 {
						continue
					}
					var temp int64 = 0
					for d := 1; d <= p; d++ {
						if p%d == 0 {
							temp += int64(mu[p/d]) * sumByPeriod[d]
						}
					}
					primitiveSum += temp
				}
				result += primitiveSum
			} else {
				result += sumByPeriod[length/2]
			}
		}

		return int(result)
	}

	return utils.Parallelise(utils.IntAcc, fn, len(ranges))
}

func Day2(input []string) []string {
	parsed := parseDay2(input)

	wg := sync.WaitGroup{}
	wg.Add(2)

	var day2Part1, day2Part2 int

	go func() {
		defer wg.Done()
		day2Part1 = solveDay2Fast(parsed, false)
	}()

	go func() {
		defer wg.Done()
		day2Part2 = solveDay2Fast(parsed, true)
	}()

	wg.Wait()
	return []string{strconv.Itoa(day2Part1), strconv.Itoa(day2Part2)}
}
