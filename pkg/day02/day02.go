package day02

import (
	"bufio"
	"strings"
)

var shapeScores = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var outcomeScores = map[string]int{
	"A X": 3,
	"A Y": 6,
	"A Z": 0,
	"B X": 0,
	"B Y": 3,
	"B Z": 6,
	"C X": 6,
	"C Y": 0,
	"C Z": 3,
}

var scoresPart2 = map[string]int{
	"A X": 0 + shapeScores["Z"],
	"A Y": 3 + shapeScores["X"],
	"A Z": 6 + shapeScores["Y"],
	"B X": 0 + shapeScores["X"],
	"B Y": 3 + shapeScores["Y"],
	"B Z": 6 + shapeScores["Z"],
	"C X": 0 + shapeScores["Y"],
	"C Y": 3 + shapeScores["Z"],
	"C Z": 6 + shapeScores["X"],
}

func Part1(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	score := 0
	for scanner.Scan() {
		// Parse moves
		line := scanner.Text()
		moves := strings.Split(line, " ")
		myMove := moves[1]

		// Shape score
		score += shapeScores[myMove]

		// Outcome score
		score += outcomeScores[line]
	}

	return score
}

func Part2(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		score += scoresPart2[line]
	}

	return score
}

func Solve(input string) (any, any) {
	return Part1(input), Part2(input)
}
