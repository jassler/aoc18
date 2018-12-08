package day02

import "fmt"

// part2
func part2(lines []string, ch chan<- string) {

	var result string
	var ok bool

all:
	for x, s1 := range lines[:len(lines)-1] {
		for _, s2 := range lines[x+1:] {
			result, ok = isSimilar(s1, s2)
			if ok {
				break all
			}
		}
	}

	ch <- fmt.Sprint("Part 2: " + result)
}

func isSimilar(s1, s2 string) (string, bool) {

	// index of the different letter
	diff := -1
	for i := range s1 {
		if s1[i] != s2[i] {
			if diff >= 0 {
				return "", false
			}

			diff = i
		}
	}

	if diff == len(s1)-1 {
		return s1[:diff], true
	}
	return s1[:diff] + s1[diff+1:], true

}
