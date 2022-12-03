package day03

import "testing"

var test_input = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

func TestPart01(t *testing.T) {
	want := 157
	got := Part1(test_input)

	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestPart02(t *testing.T) {
	want := 70
	got := Part2(test_input)

	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}
