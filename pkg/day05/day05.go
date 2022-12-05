package day05

import (
	"bufio"
	"container/list"
	"regexp"
	"strconv"
	"strings"
)

func Pop(l *list.List) any {
	if l.Len() == 0 {
		return nil
	}

	tail := l.Back()
	val := tail.Value
	l.Remove(tail)
	return val
}

func initStacks(stackString string) []*list.List {
	// Initialise stack
	re1 := regexp.MustCompile(`\d+`)
	matches := re1.FindAllString(stackString, -1)
	numStacks, _ := strconv.Atoi(matches[len(matches)-1])
	stacks := make([]*list.List, numStacks)
	for i, _ := range stacks {
		stacks[i] = list.New()
	}

	// Sort crates into stacks
	re2 := regexp.MustCompile(`[A-Z]`)
	scanner := bufio.NewScanner(strings.NewReader(stackString))
	for scanner.Scan() {
		line := scanner.Text()
		matches := re2.FindAllIndex([]byte(line), -1)

		for _, match := range matches {
			stackno := (match[1] + 3) / 4
			label := line[match[0]:match[1]]
			stacks[stackno-1].PushFront(label)
		}
	}

	return stacks
}

func topCrates(stacks []*list.List) string {
	output := ""
	for _, stack := range stacks {
		label := stack.Back().Value.(string)
		output = output + label
	}

	return output
}

func Part1(input string) string {
	// Split into stack and instructions
	inputSplit := strings.Split(input, "\n\n")
	stackString := inputSplit[0]
	instructionsString := inputSplit[1]

	stacks := initStacks(stackString)

	// Execute instructions
	re1 := regexp.MustCompile(`\d+`)
	scanner := bufio.NewScanner(strings.NewReader(instructionsString))
	for scanner.Scan() {
		line := scanner.Text()
		matches := re1.FindAllString(line, -1)

		numToMove, _ := strconv.Atoi(matches[0])
		fromStack, _ := strconv.Atoi(matches[1])
		toStack, _ := strconv.Atoi(matches[2])

		for i := 0; i < numToMove; i++ {
			label := Pop(stacks[fromStack-1])
			stacks[toStack-1].PushBack(label)
		}
	}

	return topCrates(stacks)
}

func Part2(input string) string {
	// Split into stack and instructions
	inputSplit := strings.Split(input, "\n\n")
	stackString := inputSplit[0]
	instructionsString := inputSplit[1]

	stacks := initStacks(stackString)

	// Execute instructions
	re1 := regexp.MustCompile(`\d+`)
	scanner := bufio.NewScanner(strings.NewReader(instructionsString))
	for scanner.Scan() {
		line := scanner.Text()
		matches := re1.FindAllString(line, -1)

		numToMove, _ := strconv.Atoi(matches[0])
		fromStack, _ := strconv.Atoi(matches[1])
		toStack, _ := strconv.Atoi(matches[2])

		// Pop top n crates
		var labels []string
		for i := 0; i < numToMove; i++ {
			labels = append(labels, Pop(stacks[fromStack-1]).(string))
		}

		// Push onto new stack
		for i := len(labels) - 1; i >= 0; i-- {
			stacks[toStack-1].PushBack(labels[i])
		}
	}

	return topCrates(stacks)
}

func Solve(input string) (any, any) {
	return Part1(input), Part2(input)
}
