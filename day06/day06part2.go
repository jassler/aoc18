package day06

import (
	"fmt"
)

// part2
func part2(points []point, ch chan<- string) {

	maxDistance := 10000

	// function is in part1 file
	width, height := getWidthHeight(points)
	size := 0

	// note: problem could occur if area goes beyond edges.
	// luckily it worked for me
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			if calcSummedDistance(points, point{x: x, y: y}) < maxDistance {
				size++
			}
		}
	}

	ch <- fmt.Sprintf("Part 2: Area close to multiple points is %d", size)
}

func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}

func calcSummedDistance(points []point, from point) int {
	sum := 0

	for _, p := range points {
		sum += abs(p.x-from.x) + abs(p.y-from.y)
	}

	return sum
}
