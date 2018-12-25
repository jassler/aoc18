package day12

import "fmt"

// part2
func part2(pots []int, rules []int, ch chan<- string) {

	offset := 0

	// previous plant-size
	prev := 0
	// growth between last plant-size and not
	prevDiff := 0

	// if difference has been consistent for 100 generations,
	// then I'm confident it'll stay the same in the future
	hasBeenTheSame := 0

	generation := 0
	for hasBeenTheSame < 100 {
		checkPadding(&pots, &offset)
		simulateGeneration(pots, rules)

		sum := 0
		for x, p := range pots {
			// p is either 1 or 0
			sum += p * (x + offset)
		}

		if sum-prev == prevDiff {
			hasBeenTheSame++
		} else {
			hasBeenTheSame = 0
		}
		prevDiff = sum - prev
		prev = sum
		generation++
	}

	ch <- fmt.Sprintf(
		"Part 2: After %d generations there has been a constant growth of %d, so I suspect the result after 50000000000 generations is %d",
		generation-100, prevDiff, prevDiff*(50000000000-generation)+prev)
}
