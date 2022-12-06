package day02

import "testing"

var test_input = `A Y
B X
C Z
`

func TestPart1(t *testing.T) {
	want := 15
	got := Part1(test_input)

	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestPart2(t *testing.T) {
	want := 12
	got := Part2(test_input)

	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}
