package day11

import "github.com/jassler/aoc18/fileparser"

// Start explores this problem set: https://adventofcode.com/2018/day/11
func Start(inputPath string, ch chan<- string) {
	serialNumber := fileparser.ToIntArray(inputPath)[0]
	go part1(serialNumber, ch)
	go part2(serialNumber, ch)
}
