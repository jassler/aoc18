package day09

import (
	"fmt"

	"github.com/jassler/aoc18/fileparser"
)

// Start explores this problem set: https://adventofcode.com/2018/day/9
func Start(inputPath string, ch chan<- string) {

	var players, points int
	fmt.Sscanf(fileparser.ToStringArray(inputPath)[0], "%d players; last marble is worth %d points", &players, &points)
	go part1(players, points, ch)
	go part2(players, points, ch)

}
