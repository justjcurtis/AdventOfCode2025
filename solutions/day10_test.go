package solutions

import (
	"AdventOfCode2025/utils"
	"reflect"
	"testing"
)

func TestDay10(t *testing.T) {
	input := utils.GetTestInput(10)
	expected := []string{"7", "33"}
	actual := Day10(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

func TestDay10RealData(t *testing.T) {
	input := utils.GetInputForTest(10)
	expected := []string{"522", "18105"}
	actual := Day10(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}
