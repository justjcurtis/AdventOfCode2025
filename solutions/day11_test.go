package solutions

import (
	"AdventOfCode2025/utils"
	"reflect"
	"strings"
	"testing"
)

func TestDay11Part1(t *testing.T) {
	input := utils.GetTestInput(11)
	expected := []string{"5"}
	actual := Day11(input)[:1]
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

var part2TestData = strings.Split(`svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`, "\n")

func TestDay11Part2(t *testing.T) {
	input := part2TestData
	expected := []string{"2"}
	actual := Day11(input)[1:]
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

func TestDay11RealData(t *testing.T) {
	input := utils.GetInputForTest(11)
	expected := []string{"413", "525518050323600"}
	actual := Day11(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}
