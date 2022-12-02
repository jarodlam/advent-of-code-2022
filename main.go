package main

import (
	"fmt"
	"os"

	"github.com/jarodlam/advent-of-code-2022/pkg/day01"
)

func main() {
	data, _ := os.ReadFile("input/day01.txt")
	sol1, sol2 := day01.Solve(string(data))
	fmt.Printf("Part 1: %d\nPart 2: %d\n", sol1, sol2)
}
