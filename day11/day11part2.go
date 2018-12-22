package day11

import "fmt"

// part2
func part2(serialNumber int, ch chan<- string) {

	grid := createGrid(serialNumber)

	var maxX, maxY, maxSize int
	max := 0

	// extremely inefficient way to check all 300 sizes
	for size := 1; size < 300; size++ {
		for y := 0; y < len(grid)-size; y++ {
			for x := 0; x < len(grid[y])-size; x++ {
				sum := calcPower(grid, x, y, size)
				if sum > max {
					maxX, maxY, maxSize = x+1, y+1, size
					max = sum
				}
			}
		}
	}

	ch <- fmt.Sprintf("Part 2: %d,%d,%d", maxX, maxY, maxSize)
}
