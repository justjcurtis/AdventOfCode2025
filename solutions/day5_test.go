package solutions

import (
	"AdventOfCode2025/utils"
	"reflect"
	"testing"
)

func TestDay5(t *testing.T) {
	input := utils.GetTestInput(5)
	expected := []string{"3", "14"}
	actual := Day5(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

func TestDay5RealData(t *testing.T) {
	input := utils.GetInputForTest(5)
	expected := []string{"848", "334714395325710"}
	actual := Day5(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}
