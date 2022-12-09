package day09

import (
	"testing"
)

var testInput1 = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

var testInput2 = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

func TestPart1(t *testing.T) {
	want := 13
	got := Part1(testInput1)
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestPart2(t *testing.T) {
	want1 := 1
	got1 := Part2(testInput1)
	if want1 != got1 {
		t.Errorf("Expected '%d', but got '%d'", want1, got1)
	}

	want2 := 36
	got2 := Part2(testInput2)
	if want2 != got2 {
		t.Errorf("Expected '%d', but got '%d'", want2, got2)
	}
}
