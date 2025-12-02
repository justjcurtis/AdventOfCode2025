package solutions

import (
	"AdventOfCode2025/utils"
	"reflect"
	"testing"
)

func TestDay1(t *testing.T) {
	input := utils.GetTestInput(1)
	expected := []string{"3", "6"}
	actual := Day1(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}

}

func TestDay1RealInput(t *testing.T) {
	input := utils.GetInputForTest(1)
	expected := []string{"980", "5961"}
	actual := Day1(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}
