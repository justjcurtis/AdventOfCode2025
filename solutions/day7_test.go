package solutions

import (
	"AdventOfCode2025/utils"
	"reflect"
	"testing"
)

func TestDay7(t *testing.T) {
	input := utils.GetTestInput(7)
	expected := []string{"21", "40"}
	actual := Day7(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

func TestDay7RealData(t *testing.T) {
	input := utils.GetInputForTest(7)
	expected := []string{"1533", "10733529153890"}
	actual := Day7(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}
