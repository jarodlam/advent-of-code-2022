package day12

import (
	"testing"
)

var testInput = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
`

func TestPart1(t *testing.T) {
	want := 31
	got := Part1(testInput)
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestPart2(t *testing.T) {
	want := 29
	got := Part2(testInput)
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}
