package day06

func searchForMarker(input string, markerLength int) int {
	n := len(input)
	for i := markerLength; i <= n; i++ {
		window := input[i-markerLength : i]
		set := map[rune]struct{}{}

		for _, c := range window {
			set[c] = struct{}{}
		}

		if len(set) == markerLength {
			return i
		}
	}
	panic("Failed to find marker!")
}

func Part1(input string) int {
	return searchForMarker(input, 4)
}

func Part2(input string) int {
	return searchForMarker(input, 14)
}

func Solve(input string) (any, any) {
	return Part1(input), Part2(input)
}
