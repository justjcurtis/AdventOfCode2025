package solutions

import (
	"AdventOfCode2025/utils"
	"reflect"
	"testing"
)

func TestDay8(t *testing.T) {
	input := utils.GetTestInput(8)
	expected := []string{"40", "25272"}
	actual := Day8(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

func TestDay8RealData(t *testing.T) {
	input := utils.GetInputForTest(8)
	expected := []string{"79056", "4639477"}
	actual := Day8(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}
