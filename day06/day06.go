package day06

import (
	"fmt"

	"github.com/jassler/aoc18/fileparser"
)

type point struct {
	x, y int
}

// Start explores this problem set: https://adventofcode.com/2018/day/6
func Start(inputPath string, ch chan<- string) {
	points := []point{}
	fileparser.ReadLines(inputPath, func(line string) error {

		var x, y int
		_, err := fmt.Sscanf(line, "%d, %d", &x, &y)
		points = append(points, point{x: x, y: y})
		return err

	})

	go part1(points, ch)
	go part2(points, ch)
}
