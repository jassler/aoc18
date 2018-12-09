package day05

import (
	"bytes"
	"fmt"
)

// part1
func part1(input string, ch chan<- string) {
	arr := []byte(input)
	arr = append(arr, 0)
	collapse(arr)
	ch <- fmt.Sprintf("Part 1: resulting length is %d", bytes.IndexByte(arr, 0))
}

// collapse removes all lower-upper / upper-lower-case tuples
// eg. it removes all appearances of "aA" and "Gg"
// uses C-like null-termination character
func collapse(input []byte) {
	index := 0

	for input[index+1] != 0 {
		if encodeByte(input[index]) == -encodeByte(input[index+1]) {
			i := index + 2
			for input[i] != 0 {
				input[i-2] = input[i]
				i++
			}
			input[i-2] = 0
			if index > 0 {
				index--
			}
		} else {
			index++
		}
	}
}

// encodeByte turns a letter into a number from 1 to 26 oder -1 to -26, depending on capitalization
// returns negative alphabet number if uppercase (eg. 'A' = -1, 'C' = -3, 'Z' = -26)
// returns positive alphabet number if uppercase (eg. 'a' =  1, 'c' =  3, 'z' =  26)
// prerequisit: I'm not checking, if passed byte is in alphabet
func encodeByte(b byte) int {
	if b <= 'Z' {
		// + 1 at the end, otherwise a 0 would exist
		return -(int(b-'A') + 1)
	}
	return int(b-'a') + 1
}
