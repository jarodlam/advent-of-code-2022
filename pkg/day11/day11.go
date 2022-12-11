package day11

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	Items           []int
	Operator        string
	OpValue         int
	OpOld           bool
	TestDivisibleBy int
	TrueThrow       int
	FalseThrow      int
}

func parseMonkeys(input string) []*Monkey {
	monkeys := make([]*Monkey, 0)

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := strings.Split(strings.Trim(scanner.Text(), " "), ":")
		key := line[0]

		if key == "" {
			continue
		}

		if len(key) >= 6 && key[0:6] == "Monkey" {
			monkeys = append(monkeys, &Monkey{})
			continue
		}

		value := strings.Trim(line[1], " ")
		m := monkeys[len(monkeys)-1]

		if key == "Starting items" {
			items := strings.Split(value, ", ")
			for _, item := range items {
				iInt, _ := strconv.Atoi(item)
				m.Items = append(m.Items, iInt)
			}
		}

		words := strings.Split(value, " ")

		if key == "Operation" {
			m.Operator = words[3]
			opVal := words[4]
			if opVal == "old" {
				m.OpOld = true
			} else {
				m.OpValue, _ = strconv.Atoi(opVal)
			}
		}

		if key == "Test" {
			m.TestDivisibleBy, _ = strconv.Atoi(words[2])
		}

		if key == "If true" {
			m.TrueThrow, _ = strconv.Atoi(words[3])
		}

		if key == "If false" {
			m.FalseThrow, _ = strconv.Atoi(words[3])
		}
	}

	return monkeys
}

func runSimulation(monkeys []*Monkey, rounds int, cullingFunc func(int) int) int {
	nInsp := make([]int, len(monkeys))
	for r := 1; r <= rounds; r++ {
		for i, m := range monkeys {
			for _, worry := range m.Items {
				nInsp[i]++

				// Apply operation
				var operand int
				if m.OpOld {
					operand = worry
				} else {
					operand = m.OpValue
				}
				if m.Operator == "*" {
					worry *= operand
				} else if m.Operator == "+" {
					worry += operand
				} else {
					panic("Invalid operator " + m.Operator)
				}

				// Apply division / modulo
				worry = cullingFunc(worry)

				// Throw to next monkey
				var monkeyToThrow *Monkey
				if (worry % m.TestDivisibleBy) == 0 {
					monkeyToThrow = monkeys[m.TrueThrow]
				} else {
					monkeyToThrow = monkeys[m.FalseThrow]
				}
				monkeyToThrow.Items = append(monkeyToThrow.Items, worry)
			}
			m.Items = []int{}
		}
	}

	sort.Ints(nInsp)
	return nInsp[len(nInsp)-1] * nInsp[len(nInsp)-2]
}

func Part1(input string) int {
	monkeys := parseMonkeys(input)
	return runSimulation(
		monkeys,
		20,
		func(worry int) int { return worry / 3 },
	)
}

func Part2(input string) int {
	monkeys := parseMonkeys(input)
	divisor := 1
	for _, m := range monkeys {
		divisor *= m.TestDivisibleBy
	}
	return runSimulation(
		monkeys,
		10000,
		func(worry int) int { return worry % divisor },
	)
}

func Solve(input string) (any, any) {
	return Part1(input), Part2(input)
}
