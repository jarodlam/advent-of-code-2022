package day13

import (
	"bufio"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

type result int

const (
	NOTRIGHT result = 0
	RIGHT    result = 1
	EQUAL    result = 2
)

func compareLists(pp1 []interface{}, pp2 []interface{}) result {
	i := 0
	for {
		// Check if lists have run out
		end1 := i >= len(pp1)
		end2 := i >= len(pp2)
		if end1 && end2 {
			// Both lists ran out, continue traversing packets
			return EQUAL
		} else if end1 && !end2 {
			// Left list ran out, inputs in right order
			return RIGHT
		} else if !end1 && end2 {
			// Right list ran out, inputs not in right order
			return NOTRIGHT
		}

		// Compare this element
		comp := comparePackets(pp1[i], pp2[i])
		if comp != EQUAL {
			// Found a result
			return comp
		}
		// Values are equal, continue traversing packets
		i++
	}
}

func comparePackets(p1, p2 interface{}) result {
	// fmt.Printf("Compare %v (%s) vs %v (%s)\n", p1, reflect.TypeOf(p1), p2, reflect.TypeOf(p2))
	switch pp1 := p1.(type) {
	case float64: // json parses numbers to float64
		switch pp2 := p2.(type) {
		case float64:
			// Both ints
			if pp1 < pp2 {
				return RIGHT
			} else if pp1 > pp2 {
				return NOTRIGHT
			} else {
				return EQUAL
			}
		case []interface{}:
			// Mixed types
			return comparePackets([]interface{}{pp1}, pp2)
		}
	case []interface{}:
		switch pp2 := p2.(type) {
		case float64:
			// Mixed types
			return comparePackets(pp1, []interface{}{pp2})
		case []interface{}:
			// Both lists
			return compareLists(pp1, pp2)
		}
	}

	panic("Invalid values")
}

func parsePacket(input string) interface{} {
	var j interface{}
	err := json.Unmarshal([]byte(input), &j)
	if err != nil {
		fmt.Println("Failed to parse input:", err)
	}
	return j
}

func parseInput1(input string) [][2]interface{} {
	pairs := [][2]interface{}{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		p1 := parsePacket(line)
		scanner.Scan()
		p2 := parsePacket(scanner.Text())
		p := [2]interface{}{p1, p2}
		pairs = append(pairs, p)
	}
	return pairs
}

func parseInput2(input string) []interface{} {
	packets := []interface{}{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		p := parsePacket(line)
		packets = append(packets, p)
	}
	return packets
}

func Part1(input string) int {
	pairs := parseInput1(input)
	sum := 0
	for i, pair := range pairs {
		// fmt.Println(pair)
		comp := comparePackets(pair[0], pair[1])
		switch comp {
		case RIGHT:
			sum += i + 1 // Indices should be 1-indexed
		case EQUAL:
			errmsg := fmt.Sprint("Packets found to be equal: ", pair)
			panic(errmsg)
		}
	}

	return sum
}

func Part2(input string) int {
	divider1 := []interface{}{[]interface{}{2.0}}
	divider2 := []interface{}{[]interface{}{6.0}}
	divider1str := fmt.Sprint(divider1)
	divider2str := fmt.Sprint(divider2)
	packets := parseInput2(input)
	packets = append(packets, divider1, divider2)

	sort.Slice(
		packets,
		func(i, j int) bool {
			return comparePackets(packets[i], packets[j]) == RIGHT
		},
	)

	key := 1
	for i, p := range packets {
		pstr := fmt.Sprint(p)
		if pstr == divider1str || pstr == divider2str {
			key *= i + 1
		}
	}

	return key
}

func Solve(input string) (any, any) {
	return Part1(input), Part2(input)
}
