package day16

import "github.com/jassler/aoc18/fileparser"

// Start explores this problem set: https://adventofcode.com/2018/day/16
func Start(inputPath string, ch chan<- string) {
	lines := fileparser.ToStringArray(inputPath)
	go part1(lines, ch)
	go part2(lines, ch)
}
