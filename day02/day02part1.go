package day02

import "fmt"

// part1
func part1(lines []string, ch chan<- string) {

	twos, threes := 0, 0
	for _, line := range lines {
		hasTwo, hasThree := checkLine(line)
		if hasTwo {
			twos++
		}
		if hasThree {
			threes++
		}
	}

	ch <- fmt.Sprintf("Part 1: %d * %d = %d", twos, threes, twos*threes)
}

// first return value indicates, if a letter appears twice
// second return value indicates, if a letter appears three times
func checkLine(line string) (bool, bool) {
	// kinda simulating C's nullterminating strings here
	// when counting we can shift the bytes backwards avoiding counting characters twice
	arr := []byte(fmt.Sprintf("%s%c", line, 0))
	hasTwo, hasThree := false, false

	// nullterminating character is continuously being shifted to the left
	// at some point it'll reach index 0
	for check := arr[0]; check != 0; check = arr[0] {
		count := 1
		var i int

		for i = 1; arr[i] != 0; i++ {
			if arr[i] == check {
				count++
			} else {
				arr[i-count] = arr[i]
			}
		}

		arr[i-count] = 0

		// theroetically I could check if hasTwo and hasThree is true and break
		// but meh, performance wise it probably won't make a difference
		if count == 2 {
			hasTwo = true
		}

		if count == 3 {
			hasThree = true
		}
	}

	return hasTwo, hasThree
}
