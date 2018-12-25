package day12

import (
	"strings"

	"github.com/jassler/aoc18/fileparser"
)

// Start explores this problem set: https://adventofcode.com/2018/day/12
func Start(inputPath string, ch chan<- string) {
	lines := fileparser.ToStringArray(inputPath)
	rules := make([]int, 1<<5) // 2 ^ 5 rules

	offset := len("initial state: ")
	initial := make([]int, len(lines[0])-offset)
	for x, r := range lines[0][offset:] {
		if r == '#' {
			initial[x] = 1
		} else {
			initial[x] = 0
		}
	}

	// all rules are interpreted as binary numbers,
	// where # represents a 1 and . represents a 0
	for x := 2; x < len(lines); x++ {
		rule := strings.Split(lines[x], " => ")

		// length is not 2 eg. if line is empty
		if len(rule) == 2 {
			index := 0
			for i, r := range rule[0] {
				if r == '#' {
					// 5 bits have to be filled from left to right
					index |= 1 << (4 - uint(i))
				}
			}

			// does it create a plant?
			if rule[1][0] == '#' {
				rules[index] = 1
			} else {
				rules[index] = 0
			}
		}
	}

	go part1(initial, rules, ch)
	go part2(initial, rules, ch)
}
