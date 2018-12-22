package day11

import "fmt"

// part1
func part1(serialNumber int, ch chan<- string) {

	grid := createGrid(serialNumber)

	var maxX, maxY int
	max := 0

	// look for x,y position where a 3x3 grid sums up to the largest value
	for y := 0; y < len(grid)-3; y++ {
		for x := 0; x < len(grid[y])-3; x++ {

			sum := calcPower(grid, x, y, 3)
			if sum > max {
				// index starts at 1, not 0!
				maxX, maxY = x+1, y+1
				max = sum
			}
		}
	}

	ch <- fmt.Sprintf("Part 1: %d,%d", maxX, maxY)
}

// createGrid calculates each cell in 300x300 area according to day 11 rules
func createGrid(serialNumber int) [][]int {
	grid := make([][]int, 300)
	for y := range grid {
		grid[y] = make([]int, 300)

		for x := range grid[y] {
			// set rackID to x-coord + 10
			rackID := x + 1 + 10

			// multiply rackID * y-coord
			powerLevel := rackID * (y + 1)

			// increase by grid serial number
			powerLevel += serialNumber

			// multiply with rackID
			powerLevel *= rackID

			// only hundreds digit
			powerLevel /= 100
			powerLevel -= powerLevel / 10 * 10

			// subtract 5
			powerLevel -= 5
			grid[y][x] = powerLevel
		}
	}
	return grid
}

// calcPower sums up size*size grid starting from x,y coordinate (top left corner)
func calcPower(grid [][]int, x, y, size int) int {
	sum := 0
	for relY := 0; relY < size; relY++ {
		for relX := 0; relX < size; relX++ {
			sum += grid[y+relY][x+relX]
		}
	}
	return sum
}
