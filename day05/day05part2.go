package day05

import (
	"bytes"
	"fmt"
)

// part2
func part2(input string, ch chan<- string) {

	lowestPolymerLength := len(input)
	arr := []byte(input)

	// + 1 to ensure that even if a letter never appears, a null-character can be added to the end
	toCheck := make([]byte, len(arr)+1)

	for lower, upper := byte('a'), byte('A'); lower <= 'z'; lower, upper = lower+1, upper+1 {

		// remove specified letter
		diff := 0
		var i int
		for i = 0; i < len(arr); i++ {
			if arr[i] == lower || arr[i] == upper {
				diff++
			} else {
				toCheck[i-diff] = arr[i]
			}
		}

		// null-terminating string
		toCheck[i] = 0

		// check length
		collapse(toCheck)
		var l int
		if l = bytes.IndexByte(toCheck, 0); l < lowestPolymerLength {
			lowestPolymerLength = l
		}
	}

	ch <- fmt.Sprintf("Part 2: Lowest length is %d", lowestPolymerLength)
}
