package solutions

import (
	"AdventOfCode2025/utils"
	"reflect"
	"testing"
)

func TestDay9(t *testing.T) {
	input := utils.GetTestInput(9)
	expected := []string{"50", "24"}
	actual := Day9(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

func TestDay9RealData(t *testing.T) {
	input := utils.GetInputForTest(9)
	expected := []string{"4782268188", "1574717268"}
	actual := Day9(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}
