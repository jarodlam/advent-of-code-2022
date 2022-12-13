package day12

import (
	"container/heap"
	"math"
	"strings"
)

type Pos struct {
	x int
	y int
}

type Node struct {
	pos    Pos
	weight int
	len    int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].weight < pq[j].weight }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x any) {
	item := x.(*Node)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func parseGrid(input string) ([][]int, Pos, Pos) {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	valueLookup := map[string]int{}
	for i, c := range alphabet {
		valueLookup[string(c)] = i
	}

	input = strings.Trim(input, "\n")

	grid := [][]int{}
	var start Pos
	var goal Pos
	for i, row := range strings.Split(input, "\n") {
		gridRow := []int{}
		for j, cell := range strings.Split(row, "") {
			var newVal int
			if cell == "S" {
				start = Pos{j, i}
				newVal = 0
			} else if cell == "E" {
				goal = Pos{j, i}
				newVal = 25
			} else {
				newVal = valueLookup[cell]
			}
			gridRow = append(gridRow, newVal)
		}
		grid = append(grid, gridRow)
	}

	return grid, start, goal
}

func getDist(dists map[Pos]int, pos Pos) int {
	d, ok := dists[pos]
	if !ok {
		return math.MaxInt
	}
	return d
}

func neighbours(grid [][]int, u Pos, predicate func(int) bool) []Pos {
	neigh := []Pos{}
	offsets := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	val := grid[u.y][u.x]
	for _, o := range offsets {
		newX := u.x + o[0]
		newY := u.y + o[1]
		if newX < 0 || newX >= len(grid[0]) || newY < 0 || newY >= len(grid) {
			continue
		}
		p := Pos{newX, newY}
		if predicate(grid[p.y][p.x] - val) {
			neigh = append(neigh, p)
		}
	}
	return neigh
}

func outneighbours(grid [][]int, u Pos) []Pos {
	return neighbours(grid, u, func(diff int) bool { return diff <= 1 })
}

func inneighbours(grid [][]int, u Pos) []Pos {
	return neighbours(grid, u, func(diff int) bool { return diff >= -1 })
}

func dijkstra(
	grid [][]int,
	start Pos,
	goalCondition func(Pos) bool,
	neighbourFunc func([][]int, Pos) []Pos,
) int {
	// Data structures
	pq := make(PriorityQueue, 0)
	dists := map[Pos]int{}
	parents := map[Pos]Pos{}
	visited := map[Pos]bool{}

	// Initialise at start node
	dists[start] = 0
	heap.Push(&pq, &Node{start, 0, 0})

	for pq.Len() > 0 {
		u := *(heap.Pop(&pq).(*Node))
		// fmt.Println("Visiting node", u.pos)

		if goalCondition(u.pos) {
			return u.len
		}

		if visited[u.pos] {
			continue
		}
		visited[u.pos] = true

		length := u.len + 1
		d := getDist(dists, u.pos)

		for _, v := range neighbourFunc(grid, u.pos) {
			// fmt.Println("  Expanding node", v)
			if visited[v] {
				continue
			}
			alt := d + 1
			if alt < getDist(dists, v) {
				// fmt.Println("    Reparenting (old =", getDist(dists, v), ", new = ", alt, ")")
				dists[v] = alt
				parents[v] = u.pos
				heap.Push(&pq, &Node{v, alt, length})
			}
		}
	}

	panic("Failed to find path!")
}

func Part1(input string) int {
	grid, start, goal := parseGrid(input)
	return dijkstra(
		grid,
		start,
		func(u Pos) bool { return u == goal },
		outneighbours,
	)
}

func Part2(input string) int {
	grid, _, goal := parseGrid(input)
	return dijkstra(
		grid,
		goal,
		func(u Pos) bool { return grid[u.y][u.x] == 0 },
		inneighbours,
	)
}

func Solve(input string) (any, any) {
	return Part1(input), Part2(input)
}
