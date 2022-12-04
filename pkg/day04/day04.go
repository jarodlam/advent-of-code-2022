package day04

import (
	"bufio"
	"strconv"
	"strings"
)

func Part1(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")
		range1 := strings.Split(ranges[0], "-")
		range2 := strings.Split(ranges[1], "-")

		r1a, _ := strconv.Atoi(range1[0])
		r1b, _ := strconv.Atoi(range1[1])
		r2a, _ := strconv.Atoi(range2[0])
		r2b, _ := strconv.Atoi(range2[1])

		if (r1a <= r2a) && (r2b <= r1b) || // r2 in r1
			(r2a <= r1a) && (r1b <= r2b) { // r1 in r2
			total++
		}
	}

	return total
}

func Part2(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")
		range1 := strings.Split(ranges[0], "-")
		range2 := strings.Split(ranges[1], "-")

		r1a, _ := strconv.Atoi(range1[0])
		r1b, _ := strconv.Atoi(range1[1])
		r2a, _ := strconv.Atoi(range2[0])
		r2b, _ := strconv.Atoi(range2[1])

		if (r1a <= r2a) && (r2a <= r1b) || // r2a in r1
			(r1a <= r2b) && (r2b <= r1b) || // r2b in r1
			(r2a <= r1a) && (r1a <= r2b) || // r1a in r2
			(r2a <= r1b) && (r1b <= r2b) { // r1b in r2
			total++
		}
	}

	return total
}

func Solve(input string) (int, int) {
	return Part1(input), Part2(input)
}
