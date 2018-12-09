package day03

import "fmt"

// part2
func part2(regions []*region, ch chan<- string) {

	// [elf-id] -> has not been overlapping
	isGood := make([]bool, len(regions))
	for x := range isGood {
		isGood[x] = true
	}

	var x1 int
	var r1 *region

	for x1, r1 = range regions[:len(regions)-1] {
		for x2, r2 := range regions[x1+1:] {
			if isOverlapping(r1, r2) {
				isGood[x1] = false
				isGood[x2+x1+1] = false
			}
		}

		if isGood[x1] {
			break
		}
	}

	ch <- fmt.Sprintf("Part 2: Elf #%d has the good part", regions[x1].id)
}

func isOverlapping(r1, r2 *region) bool {

	overlappingX :=
		(r1.x >= r2.x && r1.x < r2.x+r2.width) ||
			(r2.x >= r1.x && r2.x < r1.x+r1.width)

	overlappingY :=
		(r1.y >= r2.y && r1.y < r2.y+r2.height) ||
			(r2.y >= r1.y && r2.y < r1.y+r1.height)

	return overlappingX && overlappingY
}
