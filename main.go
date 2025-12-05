/*
Copyright © 2025 Jacson Curtis <justjcurtis@gmail.com>
*/
package main

import (
	"AdventOfCode2025/solutions"
	"AdventOfCode2025/utils"
	"flag"
	"fmt"
	"time"
)

type solution struct {
	day int
	fn  func([]string) []string
}

var SOLUTIONS = []solution{
	{1, solutions.Day1},
	{2, solutions.Day2},
	{3, solutions.Day3},
        {4, solutions.Day4},
        {5, solutions.Day5},
}

func main() {
	runCount := flag.Int("n", 1, "Number of times to run each solution")
	minRun := flag.Bool("min", false, "Use the minimum run time instead of the average")
	singleDay := flag.Int("d", -1, "Run only the specified day")
	readme := flag.Bool("r", false, "Generate performance for README.md file")

	flag.Parse()

	if *singleDay > len(SOLUTIONS) {
		println("Invalid day specified")
		return
	}

	if *runCount < 1 {
		*runCount = 1
	}
	if *minRun && *runCount < 2 {
		*runCount = 5000
	}

	var totalTime time.Duration
	for d, solution := range SOLUTIONS {
		if *singleDay > -1 && *singleDay != d+1 {
			continue
		}
		minElapsed := time.Duration(0)
		input := utils.GetInput(solution.day)
		start := time.Now()
		if *minRun {
			for i := 0; i < *runCount-1; i++ {
				start = time.Now()
				solution.fn(input)
				elapsed := time.Since(start)
				if elapsed < minElapsed || minElapsed == 0 {
					minElapsed = elapsed
				}
			}
			results := solution.fn(input)
			totalTime += minElapsed
			if !*readme {
				utils.PrintResults(solution.day, results)
				fmt.Printf("Day %d took %s\n", solution.day, minElapsed)
			} else {
				fmt.Printf("| Day %d | %s |\n", solution.day, utils.TruncateToDynamicUnit(minElapsed))
			}
		} else {
			for i := 0; i < *runCount-1; i++ {
				solution.fn(input)
			}
			results := solution.fn(input)
			elapsed := time.Since(start)
			totalTime += elapsed
			if !*readme {
				utils.PrintResults(solution.day, results)
				fmt.Printf("Day %d took %s\n", solution.day, elapsed/time.Duration(*runCount))
			} else {
				fmt.Printf("| Day %d | %s |\n", solution.day, utils.TruncateToDynamicUnit(elapsed/time.Duration(*runCount)))
			}

		}
		if *singleDay == -1 {
			if !*readme {
				println()
			}
		}
	}

	if *singleDay == -1 {
		if !*readme {
			println("=------ Total ------=")
			if *minRun {
				fmt.Printf("Total time: %s\n", totalTime)
			} else {
				fmt.Printf("Total time: %s\n", totalTime/time.Duration(*runCount))
			}
		} else {
			fmt.Print("| ------- | ----------------------------- |\n")
			if *minRun {
				fmt.Printf("| **Total** | **%s** |\n", utils.TruncateToDynamicUnit(totalTime))
			} else {
				fmt.Printf("| **Total** | **%s** |\n", utils.TruncateToDynamicUnit(totalTime/time.Duration(*runCount)))
			}
		}
	}
	if !*readme {
		println("=-------------------=")
	}
}
