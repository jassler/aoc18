package day05

import "github.com/jassler/aoc18/fileparser"

// Start explores this problem set: https://adventofcode.com/2018/day/5
func Start(inputPath string, ch chan<- string) {

	input := fileparser.ToStringArray(inputPath)[0]
	go part1(input, ch)
	go part2(input, ch)

}
