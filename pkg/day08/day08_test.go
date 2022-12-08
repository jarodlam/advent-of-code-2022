package day08

import (
	"testing"
)

var testInput = `30373
25512
65332
33549
35390`

func TestPart1(t *testing.T) {
	want := 21
	got := Part1(testInput)
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestPart2(t *testing.T) {
	want := 8
	got := Part2(testInput)
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}
