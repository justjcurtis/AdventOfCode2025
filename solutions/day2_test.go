package solutions

import (
	"AdventOfCode2025/utils"
	"reflect"
	"testing"
)

func TestDay2(t *testing.T) {
	input := utils.GetTestInput(2)
	expected := []string{"1227775554", "4174379265"}
	actual := Day2(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

func TestDay2RealData(t *testing.T) {
	input := utils.GetInputForTest(2)
	expected := []string{"19386344315", "34421651192"}
	actual := Day2(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}
