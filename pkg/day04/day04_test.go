package day04

import "testing"

var test_input = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

func TestPart1(t *testing.T) {
	want := 2
	got := Part1(test_input)

	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestPart2(t *testing.T) {
	want := 4
	got := Part2(test_input)

	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}
