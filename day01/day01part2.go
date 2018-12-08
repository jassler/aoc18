package day01

import "fmt"

// find first number that appears twice
func part2(numbers []int, ch chan<- string) {
	result := 0
	storedResults := map[int]bool{
		0: true,
	}

	i := 0

	// repeat calculations until first cycle
	for {
		result += numbers[i]
		if storedResults[result] {
			break
		}

		storedResults[result] = true
		i = (i + 1) % len(numbers)
	}

	ch <- fmt.Sprint("Part 2: ", result)
}
