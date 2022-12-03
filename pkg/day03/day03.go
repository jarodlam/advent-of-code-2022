package day03

import (
	"bufio"
	"strings"
)

func rucksackToSet(r string) map[rune]struct{} {
	set := map[rune]struct{}{}
	for _, c := range r {
		set[c] = struct{}{}
	}
	return set
}

func runeToPriority(r rune) int {
	// ASCII a = 97, A = 65
	// Priorities a = 1, A = 27
	if r < 97 {
		return int(r) - 38
	} else {
		return int(r) - 96
	}
}

func Part1(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	total := 0
	for scanner.Scan() {
		// Get rucksacks
		line := scanner.Text()
		r1 := line[0 : len(line)/2]
		r2 := line[len(line)/2:]

		// Convert each rucksack to set
		s1 := rucksackToSet(r1)
		s2 := rucksackToSet(r2)

		// Find the duplicate
		var found rune
		for x := range s1 {
			if _, ok := s2[x]; ok {
				found = x
				break
			}
		}

		// Add priority to total
		total += runeToPriority(found)
	}

	return total
}

func Part2(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	total := 0
	for scanner.Scan() {
		// Get three rucksacks
		r1 := scanner.Text()
		scanner.Scan()
		r2 := scanner.Text()
		scanner.Scan()
		r3 := scanner.Text()

		// Convert each rucksack to set
		s1 := rucksackToSet(r1)
		s2 := rucksackToSet(r2)
		s3 := rucksackToSet(r3)

		// Get union of s1 and s2
		var union []rune
		for x := range s1 {
			if _, ok := s2[x]; ok {
				union = append(union, x)
			}
		}

		// Find unique item in s3
		var found rune
		for _, x := range union {
			if _, ok := s3[x]; ok {
				found = x
				break
			}
		}

		// Add priority to total
		total += runeToPriority(found)
	}

	return total
}

func Solve(input string) (int, int) {
	return Part1(input), Part2(input)
}
