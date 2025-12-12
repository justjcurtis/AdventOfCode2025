package solutions

import (
	"AdventOfCode2025/utils"
	"reflect"
	"testing"
)

func TestDay12(t *testing.T) {
	input := utils.GetTestInput(12)
	expected := []string{"2"}
	actual := Day12(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

// func TestDay12RealData(t *testing.T) {
// 	input := utils.GetInputForTest(12)
// 	expected := []string{""}
// 	actual := Day12(input)
// 	if !reflect.DeepEqual(expected, actual) {
// 		t.Errorf("Expected %v but was %v", expected, actual)
// 	}
// }
