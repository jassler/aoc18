package day03

import (
	"fmt"

	"github.com/jassler/aoc18/fileparser"
)

type region struct {
	id            int
	x, y          int
	width, height int
}

var regions []*region

// Start explores this problem set: https://adventofcode.com/2018/day/3
func Start(inputPath string, ch chan<- string) {
	regions = make([]*region, 0)

	fileparser.ReadLines(inputPath, lineToRegion)

	go part1(regions, ch)
	go part2(regions, ch)
}

func lineToRegion(line string) error {
	r := region{}

	// #1 @ 509,796: 18x15
	am, err := fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &r.id, &r.x, &r.y, &r.width, &r.height)
	if err != nil {
		return fmt.Errorf("%s could not be parsed: %v", line, err)
	}
	if am != 5 {
		return fmt.Errorf("%s could not be parsed. Expected 5 numbers, got %d", line, am)
	}

	regions = append(regions, &r)
	return nil
}
