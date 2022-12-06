package day06

import (
	"testing"
)

var tests = []struct {
	input string
	want1 int
	want2 int
}{
	{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", want1: 7, want2: 19},
	{input: "bvwbjplbgvbhsrlpgdmjqwftvncz", want1: 5, want2: 23},
	{input: "nppdvjthqldpwncqszvftbrmjlhg", want1: 6, want2: 23},
	{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", want1: 10, want2: 29},
	{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", want1: 11, want2: 26},
}

func TestPart1(t *testing.T) {
	for _, tc := range tests {
		want := tc.want1
		got := Part1(tc.input)
		if want != got {
			t.Errorf("Expected '%d', but got '%d'", want, got)
		}
	}
}

func TestPart2(t *testing.T) {
	for _, tc := range tests {
		want := tc.want2
		got := Part2(tc.input)
		if want != got {
			t.Errorf("Expected '%d', but got '%d'", want, got)
		}
	}
}
