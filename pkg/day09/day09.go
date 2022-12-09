package day09

import (
	"bufio"
	"strconv"
	"strings"
)

type GridPos struct {
	x int
	y int
}

func dirToDelta(dir string) GridPos {
	switch dir {
	case "U":
		return GridPos{0, 1}
	case "D":
		return GridPos{0, -1}
	case "L":
		return GridPos{-1, 0}
	case "R":
		return GridPos{1, 0}
	default:
		panic("Invalid direction '" + dir + "'")
	}
}

func applyDelta(pos GridPos, delta GridPos) GridPos {
	return GridPos{pos.x + delta.x, pos.y + delta.y}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getTailDelta(H GridPos, T GridPos) GridPos {
	dx := H.x - T.x
	dy := H.y - T.y

	// Check if diagonal + 1
	threshold := 1
	if abs(dx) > 0 && abs(dy) > 0 && (abs(dx)+abs(dy)) > 2 {
		threshold = 0
	}

	f := func(d int) int {
		if d > threshold {
			return 1
		} else if d < -threshold {
			return -1
		} else {
			return 0
		}
	}
	return GridPos{f(dx), f(dy)}
}

func Part1(input string) int {
	H := GridPos{0, 0}
	T := GridPos{0, 0}
	visited := map[GridPos]struct{}{}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		// Parse instruction
		line := strings.Split(scanner.Text(), " ")
		dir := line[0]
		steps, _ := strconv.Atoi(line[1])

		// Apply instruction
		deltaH := dirToDelta(dir)
		for i := 0; i < steps; i++ {
			// Move H
			H = applyDelta(H, deltaH)

			// Move T
			deltaT := getTailDelta(H, T)
			T = applyDelta(T, deltaT)

			visited[T] = struct{}{}
		}
	}

	return len(visited)
}

func Part2(input string) int {
	ropeLen := 10
	rope := make([]GridPos, ropeLen)
	for i := 0; i < ropeLen; i++ {
		rope = append(rope, GridPos{0, 0})
	}
	visited := map[GridPos]struct{}{}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		// Parse instruction
		line := strings.Split(scanner.Text(), " ")
		dir := line[0]
		steps, _ := strconv.Atoi(line[1])

		// Apply instruction
		deltaH := dirToDelta(dir)
		for i := 0; i < steps; i++ {
			// Move head
			rope[0] = applyDelta(rope[0], deltaH)

			// Move all knots one-by-one
			for j := 1; j < ropeLen; j++ {
				deltaT := getTailDelta(rope[j-1], rope[j])
				rope[j] = applyDelta(rope[j], deltaT)
			}

			// Record tail
			visited[rope[ropeLen-1]] = struct{}{}
		}
	}

	return len(visited)
}

func Solve(input string) (any, any) {
	return Part1(input), Part2(input)
}
