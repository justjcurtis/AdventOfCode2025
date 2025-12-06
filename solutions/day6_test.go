package solutions

import (
	"AdventOfCode2025/utils"
	"reflect"
	"testing"
)

func TestDay6(t *testing.T) {
	input := utils.GetTestInput(6)
	expected := []string{"4277556", "3263827"}
	actual := Day6(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

func TestDay6RealData(t *testing.T) {
	input := utils.GetInputForTest(6)
	expected := []string{"4309240495780", "9170286552289"}
	actual := Day6(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}
