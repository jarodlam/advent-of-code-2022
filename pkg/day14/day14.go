package day14

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const (
	MAX_X = 1000
	MAX_Y = 200
	SRC_X = 500
	SRC_Y = 0
)

func parseCoords(line string) [][2]int {
	coordStr := strings.Split(line, " -> ")
	coords := make([][2]int, len(coordStr))
	for i, str := range coordStr {
		split := strings.Split(str, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		coords[i] = [2]int{x, y}
	}
	return coords
}

func drawLine(canvas *([][]rune), c1, c2 [2]int) {
	if c1[1] == c2[1] {
		// Horizontal line
		y := c1[1]
		x1, x2 := c1[0], c2[0]
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		for x := x1; x <= x2; x++ {
			(*canvas)[y][x] = '#'
		}
	} else if c1[0] == c2[0] {
		// Vertical line
		x := c1[0]
		y1, y2 := c1[1], c2[1]
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for y := y1; y <= y2; y++ {
			(*canvas)[y][x] = '#'
		}
	} else {
		errmsg := fmt.Sprint("Coordinates are not aligned along an axis:", c1, c2)
		panic(errmsg)
	}
}

func parseInput(input string, addFloor bool) [][]rune {
	// Init canvas
	canvas := make([][]rune, MAX_Y)
	blankRow := make([]rune, MAX_X)
	for i := range blankRow {
		blankRow[i] = ' '
	}
	for i := range canvas {
		canvas[i] = append(canvas[i], blankRow...)
	}

	maxRockY := 0

	// Execute instructions
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		coords := parseCoords(scanner.Text())
		if addFloor && coords[0][1] > maxRockY {
			maxRockY = coords[0][1]
		}
		for i := 1; i < len(coords); i++ {
			drawLine(&canvas, coords[i-1], coords[i])
			if addFloor && coords[i][1] > maxRockY {
				maxRockY = coords[i][1]
			}
		}
	}

	// Add floor
	if addFloor {
		maxRockY += 2
		for i := range canvas[maxRockY] {
			canvas[maxRockY][i] = '#'
		}
	}

	return canvas
}

func printCanvas(canvas [][]rune, minX int) {
	for _, ln := range canvas {
		for x, r := range ln {
			if x < minX {
				continue
			}
			fmt.Print(string(r))
		}
		fmt.Println()
	}
}

func Part1(input string) int {
	canvas := parseInput(input, false)

	// Simulate
	for count := 0; ; count++ {
		sandX := SRC_X
		sandY := SRC_Y

		// Apply sand rules until resting or void
		for {
			if canvas[sandY+1][sandX] == ' ' {
				// A unit of sand always falls down one step if possible.
				sandY += 1
			} else if canvas[sandY+1][sandX-1] == ' ' {
				// If the tile immediately below is blocked (by rock or sand), the unit of
				// sand attempts to instead move diagonally one step down and to the left.
				sandY += 1
				sandX -= 1
			} else if canvas[sandY+1][sandX+1] == ' ' {
				// If that tile is blocked, the unit of sand attempts to instead move
				// diagonally one step down and to the right.
				sandY += 1
				sandX += 1
			} else {
				// If all three possible destinations are blocked, the unit of sand comes
				// to rest and no longer moves.
				break
			}

			if sandY >= (MAX_Y - 1) {
				// Sand fell into the void
				//printCanvas(canvas, 450)
				return count
			}
		}

		// Paint sand to canvas
		canvas[sandY][sandX] = 'o'
	}
}

func Part2(input string) int {
	canvas := parseInput(input, true)

	// Simulate
	for count := 1; ; count++ {
		sandX := SRC_X
		sandY := SRC_Y

		// Apply sand rules until resting or void
		for {
			if canvas[sandY+1][sandX] == ' ' {
				// A unit of sand always falls down one step if possible.
				sandY += 1
			} else if canvas[sandY+1][sandX-1] == ' ' {
				// If the tile immediately below is blocked (by rock or sand), the unit of
				// sand attempts to instead move diagonally one step down and to the left.
				sandY += 1
				sandX -= 1
			} else if canvas[sandY+1][sandX+1] == ' ' {
				// If that tile is blocked, the unit of sand attempts to instead move
				// diagonally one step down and to the right.
				sandY += 1
				sandX += 1
			} else {
				// If all three possible destinations are blocked, the unit of sand comes
				// to rest and no longer moves.
				if sandX == SRC_X && sandY == SRC_Y {
					// Simulate until source is blocked.
					// printCanvas(canvas, 450)
					return count
				}
				break
			}
		}

		// Paint sand to canvas
		canvas[sandY][sandX] = 'o'
	}
}

func Solve(input string) (any, any) {
	return Part1(input), Part2(input)
}
