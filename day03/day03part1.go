package day03

import "fmt"

// part1
// count overlapping regions
func part1(regions []*region, ch chan<- string) {

	fabric := [1000][1000]byte{}
	count := 0

	// iterate region
	for _, r := range regions {
		// iterate rows
		for y := r.y; y < r.height+r.y; y++ {
			// iterate columns
			for x := r.x; x < r.width+r.x; x++ {

				// as long as it's < 2, it hasn't been overwritten by someone already
				if fabric[y][x] < 2 {
					fabric[y][x]++
					if fabric[y][x] == 2 {
						count++
					}
				}
			}
		}
	}

	ch <- fmt.Sprintf("Part 1: %d", count)
}
