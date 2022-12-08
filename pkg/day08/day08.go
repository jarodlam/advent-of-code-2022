package day08

import (
	"fmt"
	"strconv"
	"strings"
)

func scenicScore(y int, x int, forest [][]int, nrow int, ncol int) int {
	scenicScore := 1
	maxHeight := forest[y][x]
	var dist int

	// Look up
	dist = 0
	for i := y - 1; i >= 0; i-- {
		val := forest[i][x]
		dist++
		if val >= maxHeight {
			break
		}
	}
	fmt.Printf("%d ", dist)
	if dist == 0 {
		return 0
	}
	scenicScore *= dist

	// Look left
	dist = 0
	for i := x - 1; i >= 0; i-- {
		val := forest[y][i]
		dist++
		if val >= maxHeight {
			break
		}
	}
	fmt.Printf("%d ", dist)
	if dist == 0 {
		return 0
	}
	scenicScore *= dist

	// Look down
	dist = 0
	for i := y + 1; i < nrow; i++ {
		val := forest[i][x]
		dist++
		if val >= maxHeight {
			break
		}
	}
	fmt.Printf("%d ", dist)
	if dist == 0 {
		return 0
	}
	scenicScore *= dist

	// Look right
	dist = 0
	for i := x + 1; i < ncol; i++ {
		val := forest[y][i]
		dist++
		if val >= maxHeight {
			break
		}
	}
	fmt.Printf("%d ", dist)
	if dist == 0 {
		return 0
	}
	scenicScore *= dist

	return scenicScore
}

func Part1(input string) int {
	forest := make([][]int, 0)
	seen := make([][]bool, 0)
	totalSeen := 0

	// Populate forest
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

	nrow := len(forest)
	ncol := len(forest[0])

	for i := 0; i < nrow; i++ {
		// Look at left side
		maxHeight := -1
		for j := 0; j < ncol; j++ {
			if forest[i][j] > maxHeight {
				// Tree is visible
				maxHeight = forest[i][j]
				if !seen[i][j] {
					seen[i][j] = true
					totalSeen++
				}
			}
		}

		// Look at right side
		maxHeight = -1
		for j := ncol - 1; j >= 0; j-- {
			if forest[i][j] > maxHeight {
				// Tree is visible
				maxHeight = forest[i][j]
				if !seen[i][j] {
					seen[i][j] = true
					totalSeen++
				}
			}
		}
	}

	for j := 0; j < ncol; j++ {
		// Look at top side
		maxHeight := -1
		for i := 0; i < nrow; i++ {
			if forest[i][j] > maxHeight {
				// Tree is visible
				maxHeight = forest[i][j]
				if !seen[i][j] {
					seen[i][j] = true
					totalSeen++
				}
			}
		}

		// Look at right side
		maxHeight = -1
		for i := nrow - 1; i >= 0; i-- {
			if forest[i][j] > maxHeight {
				// Tree is visible
				maxHeight = forest[i][j]
				if !seen[i][j] {
					seen[i][j] = true
					totalSeen++
				}
			}
		}
	}

	return totalSeen
}

func Part2(input string) int {
	forest := make([][]int, 0)

	// Populate forest
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
	}

	nrow := len(forest)
	ncol := len(forest[0])

	// Scan trees
	maxScore := 0
	for i := 0; i < nrow; i++ {
		for j := 0; j < ncol; j++ {
			fmt.Printf("(%d, %d) %d: ", i, j, forest[i][j])
			score := scenicScore(i, j, forest, nrow, ncol)
			fmt.Printf(", %d\n", score)
			if score > maxScore {
				maxScore = score
			}
		}
		fmt.Println()
	}

	return maxScore
}

func Solve(input string) (any, any) {
	return Part1(input), Part2(input)
}
