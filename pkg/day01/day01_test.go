package day01

import "testing"

var test_input = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

func TestPart01(t *testing.T) {
	want := 24000
	got := Part1(test_input)

	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}
