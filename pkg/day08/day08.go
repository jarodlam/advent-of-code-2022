package day08

import (
	"strconv"
	"strings"
)

func populateForest(input string) ([][]int, [][]bool) {
	forest := make([][]int, 0)
	seen := make([][]bool, 0)

	for _, row := range strings.Split(input, "\n") {
		if len(row) == 0 {
			continue
		}
		rowSplit := strings.Split(row, "")
		newRow := make([]int, 0)
		for _, c := range rowSplit {
			integer, _ := strconv.Atoi(c)
			newRow = append(newRow, integer)
		}
		forest = append(forest, newRow)
		seen = append(seen, make([]bool, len(newRow)))
	}

	return forest, seen
}

func scenicScore(y int, x int, forest [][]int, nrow int, ncol int) int {
	scenicScore := 1
	maxHeight := forest[y][x]
	lookInDirection := func(start int, predicate func(int) bool, inc int, val func(int) int) int {
		dist := 0
		for i := start; predicate(i); i += inc {
			dist++
			if val(i) >= maxHeight {
				break
			}
		}
		return dist
	}
	predDec := func(i int) bool { return i >= 0 }
	predIncRow := func(i int) bool { return i < nrow }
	predIncCol := func(i int) bool { return i < ncol }
	valX := func(i int) int { return forest[i][x] }
	valY := func(i int) int { return forest[y][i] }

	scenicScore *= lookInDirection(y-1, predDec, -1, valX)   // up
	scenicScore *= lookInDirection(x-1, predDec, -1, valY)   // left
	scenicScore *= lookInDirection(y+1, predIncRow, 1, valX) // down
	scenicScore *= lookInDirection(x+1, predIncCol, 1, valY) // right

	return scenicScore
}

func Part1(input string) int {
	forest, seen := populateForest(input)
	nrow := len(forest)
	ncol := len(forest[0])
	totalSeen := 0

	checkTree := func(i int, j int, maxHeight int, cols bool) int {
		if cols {
			tmp := i
			i = j
			j = tmp
		}
		if forest[i][j] > maxHeight {
			if !seen[i][j] {
				seen[i][j] = true
				totalSeen++
			}
			return forest[i][j]
		}
		return maxHeight
	}

	lookInDirection := func(outerLim int, innerLim int, cols bool) {
		for i := 0; i < outerLim; i++ {
			maxHeight := -1
			for j := 0; j < innerLim; j++ {
				maxHeight = checkTree(i, j, maxHeight, cols)
			}
			maxHeight = -1
			for j := innerLim - 1; j >= 0; j-- {
				maxHeight = checkTree(i, j, maxHeight, cols)
			}
		}
	}

	lookInDirection(nrow, ncol, false) // left/right
	lookInDirection(ncol, nrow, true)  // top/bottom

	return totalSeen
}

func Part2(input string) int {
	forest, _ := populateForest(input)
	nrow := len(forest)
	ncol := len(forest[0])

	// Scan each tree
	maxScore := 0
	for i := 0; i < nrow; i++ {
		for j := 0; j < ncol; j++ {
			score := scenicScore(i, j, forest, nrow, ncol)
			if score > maxScore {
				maxScore = score
			}
		}
	}

	return maxScore
}

func Solve(input string) (any, any) {
	return Part1(input), Part2(input)
}
