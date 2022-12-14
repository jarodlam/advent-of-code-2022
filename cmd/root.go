/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/jarodlam/advent-of-code-2022/pkg/day01"
	"github.com/jarodlam/advent-of-code-2022/pkg/day02"
	"github.com/jarodlam/advent-of-code-2022/pkg/day03"
	"github.com/jarodlam/advent-of-code-2022/pkg/day04"
	"github.com/jarodlam/advent-of-code-2022/pkg/day05"
	"github.com/jarodlam/advent-of-code-2022/pkg/day06"
	"github.com/jarodlam/advent-of-code-2022/pkg/day07"
	"github.com/jarodlam/advent-of-code-2022/pkg/day08"
	"github.com/jarodlam/advent-of-code-2022/pkg/day09"
	"github.com/jarodlam/advent-of-code-2022/pkg/day10"
	"github.com/jarodlam/advent-of-code-2022/pkg/day11"
	"github.com/jarodlam/advent-of-code-2022/pkg/day12"
	"github.com/jarodlam/advent-of-code-2022/pkg/day13"
	"github.com/jarodlam/advent-of-code-2022/pkg/day14"
	"github.com/jarodlam/advent-of-code-2022/pkg/day15"
	"github.com/spf13/cobra"
)

var Input string

// Solution functions to run for each day
var dayFunctions = map[string]func(string) (any, any){
	"day01": day01.Solve,
	"day02": day02.Solve,
	"day03": day03.Solve,
	"day04": day04.Solve,
	"day05": day05.Solve,
	"day06": day06.Solve,
	"day07": day07.Solve,
	"day08": day08.Solve,
	"day09": day09.Solve,
	"day10": day10.Solve,
	"day11": day11.Solve,
	"day12": day12.Solve,
	"day13": day13.Solve,
	"day14": day14.Solve,
	"day15": day15.Solve,
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aoc2022 day",
	Short: "Run a solution for Advent of Code 2022.",
	Long:  "Solutions for Advent of Code 2022.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("must provide day argument e.g. day01")
		}

		return runSolution((args[0]))
	},
}

func runSolution(day string) error {
	// Get solution function
	solveFunc, ok := dayFunctions[day]
	if !ok {
		return errors.New("invalid day " + day)
	}

	// Default input file to input/<day>.txt
	if Input == "" {
		Input = path.Join("input", day+".txt")
	}

	// Read file
	data, err := os.ReadFile(Input)
	if err != nil {
		return err
	}

	// Run solution
	sol1, sol2 := solveFunc(string(data))
	fmt.Println("Part 1:", sol1)
	fmt.Println("Part 2:", sol2)

	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&Input, "input", "i", "", "Input file path")
}
