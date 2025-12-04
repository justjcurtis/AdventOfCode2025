package solutions

import (
	"AdventOfCode2025/utils"
	"reflect"
	"testing"
)

func TestDay4(t *testing.T) {
	input := utils.GetTestInput(4)
	expected := []string{"13", "43"}
	actual := Day4(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

func TestDay4RealData(t *testing.T) {
	input := utils.GetInputForTest(4)
	expected := []string{"1569", "9280"}
	actual := Day4(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}
