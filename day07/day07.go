package day07

import (
	"fmt"

	"github.com/jassler/aoc18/fileparser"
)

type instructionTuple struct {
	first rune
	then  rune
}

// Start explores this problem set: https://adventofcode.com/2018/day/7
func Start(inputPath string, ch chan<- string) {
	tuples := []*instructionTuple{}

	// parse all steps
	fileparser.ReadLines(inputPath, func(line string) error {
		var parent, child rune
		_, err := fmt.Sscanf(line, "Step %c must be finished before step %c can begin.", &parent, &child)
		tuples = append(tuples, &instructionTuple{first: parent, then: child})
		return err
	})

	go part1(tuples, ch)
	go part2(tuples, ch)
}
