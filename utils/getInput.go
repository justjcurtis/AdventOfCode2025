package utils

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

func parseInput(f *os.File) []string {
	lines := make([]string, 0)
	rd := bufio.NewReader(f)
	for {
		line, _, err := rd.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		lines = append(lines, string(line))
	}
	return lines
}

func GetTestInput(day int) []string {
	cwd, _ := os.Getwd()
	parentDir := filepath.Dir(cwd)
	f, err := os.Open(parentDir + "/puzzleInput/test_" + strconv.Itoa(day) + ".txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return parseInput(f)
}

func GetInputForTest(day int) []string {
	cwd, _ := os.Getwd()
	parentDir := filepath.Dir(cwd)
	f, err := os.Open(parentDir + "/puzzleInput/day_" + strconv.Itoa(day) + ".txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return parseInput(f)
}

func GetInput(day int) []string {
	cwd, _ := os.Getwd()
	f, err := os.Open(cwd + "/puzzleInput/day_" + strconv.Itoa(day) + ".txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return parseInput(f)
}
