package day15

import (
	"bufio"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Sensor struct {
	sX   int // Sensor x
	sY   int // Sensor y
	bX   int // Beacon x
	bY   int // Beacon y
	d    int // Manhattan distance from sensor to beacon
	minX int // Bbox min x
	minY int // Bbox min y
	maxX int // Bbox max x
	maxY int // Bbox max y
}

func abs(x int) int { return int(math.Abs(float64(x))) }

func parseInput(input string) []Sensor {
	sensors := []Sensor{}
	re := regexp.MustCompile(`-?\d+`)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		coords := re.FindAllString(scanner.Text(), 4)
		sX, _ := strconv.Atoi(coords[0])
		sY, _ := strconv.Atoi(coords[1])
		bX, _ := strconv.Atoi(coords[2])
		bY, _ := strconv.Atoi(coords[3])
		d := abs(bX-sX) + abs(bY-sY)
		minX, minY := sX-d, sY-d
		maxX, maxY := sX+d, sY+d

		s := Sensor{sX, sY, bX, bY, d, minX, minY, maxX, maxY}
		sensors = append(sensors, s)
	}
	return sensors
}

func rowCollision(sensors []Sensor, rowToCheck int, maxDim int) ([][2]int, int) {
	// Check collision against each sensor and beacon
	intervals := [][2]int{}
	beacons := map[[2]int]struct{}{}
	for _, s := range sensors {
		// Check bounding box
		if !(s.minY <= rowToCheck && rowToCheck <= s.maxY) {
			continue
		}

		// Check for beacons
		if s.bY == rowToCheck {
			beacons[[2]int{s.bX, s.bY}] = struct{}{}
		}

		// Find interval covered
		d_y := abs(s.sY - rowToCheck)
		d_x := s.d - d_y
		intvl := [2]int{s.sX - d_x, s.sX + d_x}

		// Check maxDim
		if intvl[0] > maxDim {
			// Entirely outside range, skip
			continue
		} else if intvl[1] > maxDim {
			// Half outside range, truncate
			intvl[1] = maxDim
		}

		intervals = append(intervals, intvl)
	}

	// Merge intervals
	sort.Slice(
		intervals,
		func(i, j int) bool { return intervals[i][0] < intervals[j][0] },
	)
	intervalsMerged := [][2]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		// Check previous interval if overlaps with current
		prev := intervalsMerged[len(intervalsMerged)-1]
		curr := intervals[i]
		if curr[1] <= prev[1] {
			// Fully enclosed, skip
			continue
		} else if curr[0] <= prev[1] {
			// Partially enclosed, extend
			intervalsMerged[len(intervalsMerged)-1][1] = curr[1]
		} else {
			// Not enclosed, make new
			intervalsMerged = append(intervalsMerged, curr)
		}
	}

	return intervalsMerged, len(beacons)
}

func Part1(input string, rowToCheck int) int {
	sensors := parseInput(input)
	intervals, nBeacons := rowCollision(sensors, rowToCheck, math.MaxInt)

	// Sum intervals and subtract beacons
	sum := -nBeacons
	for _, intvl := range intervals {
		sum += intvl[1] - intvl[0] + 1 // Inclusive of both bounds so add 1
	}

	return sum
}

func Part2(input string, maxDim int) int {
	sensors := parseInput(input)

	for row := 0; row <= maxDim; row++ {
		intervals, _ := rowCollision(sensors, row, maxDim)
		if len(intervals) > 1 {
			return (intervals[0][1]+1)*4000000 + row
		}
	}

	panic("Failed to find beacon!")
}

func Solve(input string) (any, any) {
	return Part1(input, 2000000), Part2(input, 4000000)
}
