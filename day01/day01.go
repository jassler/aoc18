package day01

import "github.com/jassler/aoc18/fileparser"

// Start explores this problem set: https://adventofcode.com/2018/day/1
func Start(inputPath string, ch chan<- string) {
	numbers := fileparser.ToIntArray(inputPath)
	go part1(numbers, ch)
	go part2(numbers, ch)
}
