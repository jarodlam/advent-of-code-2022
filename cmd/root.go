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
	"github.com/spf13/cobra"
)

var Input string

var dayFunctions = map[string]func(string) (int, int){
	"day01": day01.Solve,
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
	sol1, sol2 := dayFunctions[day](string(data))
	fmt.Printf("Part 1: %d\nPart 2: %d\n", sol1, sol2)

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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.advent-of-code-2022.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&Input, "input", "i", "", "Input file path")
}
