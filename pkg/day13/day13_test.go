package day13

import (
	"testing"
)

var testInput = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
`

func TestPart1(t *testing.T) {
	want := 13
	got := Part1(testInput)
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestPart2(t *testing.T) {
	want := 140
	got := Part2(testInput)
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}
