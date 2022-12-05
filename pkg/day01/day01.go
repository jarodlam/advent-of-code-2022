package day01

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
)

func Part1(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	largest := 0
	current_total := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if current_total > largest {
				largest = current_total
			}
			current_total = 0
			continue
		}

		i, _ := strconv.Atoi(line)
		current_total += i
	}

	return largest
}

func Part2(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	totals := make([]int, 0)
	current_total := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			totals = append(totals, current_total)
			current_total = 0
			continue
		}

		i, _ := strconv.Atoi(line)
		current_total += i
	}
	// Add the last total
	if current_total != 0 {
		totals = append(totals, current_total)
	}

	// Get top 3
	sort.Sort(sort.Reverse(sort.IntSlice(totals)))
	top3 := totals[0] + totals[1] + totals[2]

	return top3
}

func Solve(input string) (any, any) {
	return Part1(input), Part2(input)
}
