package day05

import "testing"

var test_input = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

func TestPart01(t *testing.T) {
	want := "CMZ"
	got := Part1(test_input)

	if want != got {
		t.Errorf("Expected '%s', but got '%s'", want, got)
	}
}

func TestPart02(t *testing.T) {
	want := "MCD"
	got := Part2(test_input)

	if want != got {
		t.Errorf("Expected '%s', but got '%s'", want, got)
	}
}
