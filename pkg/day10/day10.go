package day10

import (
	"bufio"
	"strconv"
	"strings"
)

func Part1(input string) int {
	cycle := 0
	X := 1
	total := 0
	isInteresting := func(cycle int) bool {
		return ((cycle-20)%40) == 0 && (cycle <= 220)
	}
	incrementCycle := func() {
		cycle++
		if isInteresting(cycle) {
			total += X * cycle
		}
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		command := line[0]

		// noop
		if command == "noop" {
			incrementCycle()
			continue
		}

		// addx
		incrementCycle()
		incrementCycle()
		V, _ := strconv.Atoi(line[1])
		X += V
	}

	return total
}

func Part2(input string) string {
	beamPos := 0
	spritePos := 1
	output := "\n"

	incrementCycle := func() {
		if beamPos >= (spritePos-1) && (spritePos+1) >= beamPos {
			output += "#"
		} else {
			output += "."
		}
		beamPos++
		if beamPos >= 40 {
			output += "\n"
			beamPos = 0
		}
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		command := line[0]

		// noop
		if command == "noop" {
			incrementCycle()
			continue
		}

		// addx
		incrementCycle()
		incrementCycle()
		V, _ := strconv.Atoi(line[1])
		spritePos += V
	}

	return strings.TrimRight(output, "\n")
}

func Solve(input string) (any, any) {
	return Part1(input), Part2(input)
}
