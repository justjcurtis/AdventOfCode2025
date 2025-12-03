package solutions

import (
	"AdventOfCode2025/utils"
	"reflect"
	"testing"
)

func TestDay3(t *testing.T) {
	input := utils.GetTestInput(3)
	expected := []string{"357", "3121910778619"}
	actual := Day3(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

func TestDay3RealData(t *testing.T) {
	input := utils.GetInputForTest(3)
	expected := []string{"16946", "168627047606506"}
	actual := Day3(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}
