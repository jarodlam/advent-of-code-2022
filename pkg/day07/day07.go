package day07

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Directory struct {
	Name    string
	Files   map[string]int
	SubDirs map[string]*Directory
	Parent  *Directory
}

func mkdir(name string, parent *Directory) *Directory {
	return &Directory{
		name,
		map[string]int{},
		map[string]*Directory{},
		parent,
	}
}

func processCommand(line string, cwd *Directory, rootDir *Directory, scanner *bufio.Scanner) *Directory {
	args := strings.Split(line, " ")
	command := args[1]

	// Parse command
	if command == "cd" {
		cdDir := args[2]

		// Handle special dirs
		if cdDir == ".." {
			cwd = (*cwd).Parent
			return cwd
		} else if cdDir == "/" {
			cwd = rootDir
			return cwd
		}

		// Check if dir exists
		// _, ok := (*cwd).SubDirs[cdDir]
		// if !ok {
		// 	(*cwd).SubDirs[cdDir] = mkdir(cdDir, cwd)
		// }

		// Change to this directory
		cwd = (*cwd).SubDirs[cdDir]

	} else if command == "ls" {
		for scanner.Scan() {
			lsLine := scanner.Text()

			// Reached end of ls output, process next command
			if string(lsLine[0]) == "$" {
				return processCommand(lsLine, cwd, rootDir, scanner)
			}

			lsArgs := strings.Split(lsLine, " ")

			if lsArgs[0] == "dir" {
				// Add directory
				(*cwd).SubDirs[lsArgs[1]] = mkdir(lsArgs[1], cwd)
			} else {
				// Add file
				size, _ := strconv.Atoi(lsArgs[0])
				cwd.Files[lsArgs[1]] = size
			}
		}
	}

	return cwd
}

func indentFileTree(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print("  ")
	}
}

func printFileTree(dir *Directory, depth int) {
	indentFileTree(depth)
	fmt.Printf("- %s (dir)\n", dir.Name)

	for _, d := range dir.SubDirs {
		printFileTree(d, depth+1)
	}

	for name, sz := range dir.Files {
		indentFileTree(depth + 1)
		fmt.Printf("- %s (file, size=%d)\n", name, sz)
	}
}

func depthFirstSearch1(dir *Directory, matchingSize *int) int {
	dirSize := 0
	for _, dir := range dir.SubDirs {
		dirSize += depthFirstSearch1(dir, matchingSize)
	}
	for _, fileSize := range dir.Files {
		dirSize += fileSize
	}

	if dirSize <= 100000 {
		*matchingSize += dirSize
	}

	return dirSize
}

func depthFirstSearch2(dir *Directory, dirSizes *[]int) int {
	dirSize := 0
	for _, dir := range dir.SubDirs {
		dirSize += depthFirstSearch2(dir, dirSizes)
	}
	for _, fileSize := range dir.Files {
		dirSize += fileSize
	}

	*dirSizes = append(*dirSizes, dirSize)

	return dirSize
}

func Part1(input string) int {
	rootDir := mkdir("/", nil)
	cwd := rootDir

	// Populate tree
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		cwd = processCommand(line, cwd, rootDir, scanner)
	}

	// Traverse tree
	matchingSize := 0
	depthFirstSearch1(rootDir, &matchingSize)

	//printFileTree(rootDir, 0)
	return matchingSize
}

func Part2(input string) int {
	rootDir := mkdir("/", nil)
	cwd := rootDir

	// Populate tree
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		cwd = processCommand(line, cwd, rootDir, scanner)
	}

	// Traverse tree
	dirSizes := make([]int, 0)
	usedSpace := depthFirstSearch2(rootDir, &dirSizes)

	// Find directory to delete
	freeSpace := 70000000 - usedSpace
	spaceToDelete := 30000000 - freeSpace
	sort.Ints(dirSizes)
	for _, x := range dirSizes {
		if x >= spaceToDelete {
			return x
		}
	}

	panic("Failed to find directory to delete!")
}

func Solve(input string) (any, any) {
	return Part1(input), Part2(input)
}
