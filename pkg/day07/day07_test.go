package day07

import (
	"testing"
)

var testInput = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func TestPart1(t *testing.T) {
	want := 95437
	got := Part1(testInput)
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestPart2(t *testing.T) {
	want := 24933642
	got := Part2(testInput)
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}
