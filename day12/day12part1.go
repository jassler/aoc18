package day12

import (
	"fmt"
)

// part1
func part1(pots []int, rules []int, ch chan<- string) {

	offset := 0
	for generation := 0; generation < 20; generation++ {
		checkPadding(&pots, &offset)
		simulateGeneration(pots, rules)
	}

	sum := 0
	for x, p := range pots {
		// p is either 1 or 0
		sum += p * (x + offset)
	}

	ch <- fmt.Sprintf("Part 1: %d (%s)", sum, createPotsString(pots))
}

// checkPadding makes sure there's always two empty spaces to the left-most
// and right-most plant (eg. if pots array is .#..# -> ..#..#..)
func checkPadding(pots *[]int, offset *int) {

	if (*pots)[0] == 1 {
		// add two to the left
		*pots = append([]int{0, 0}, (*pots)...)
		*offset -= 2
	} else if (*pots)[1] == 1 {
		// add one to the left
		*pots = append([]int{0}, (*pots)...)
		*offset--
	}

	l := len(*pots)
	if (*pots)[l-1] == 1 {
		// add two to the right
		*pots = append(*pots, 0, 0)
	} else if (*pots)[l-2] == 1 {
		// add one to the right
		*pots = append(*pots, 0)
	}
}

// createPotsString turns pots array into a string
// where each pot is either mapped to a '#' (contains plant) or '.' (no plant)
func createPotsString(pots []int) string {
	res := make([]rune, len(pots))
	for i, p := range pots {
		if p == 1 {
			res[i] = '#'
		} else {
			res[i] = '.'
		}
	}
	return string(res)
}

// evaluate checks, if with current constellation a new plant appears (1) or not (0)
// shift should be 4 by default (eg. if provided pots slice has 5 elements)
func evaluate(pots []int, rules []int, shift uint) int {
	index := 0
	for i, p := range pots {
		index |= p << (shift - uint(i))
	}
	return rules[index]
}

func simulateGeneration(pots []int, rules []int) {

	// evaluate special cases, last and first two indeces
	l := len(pots)
	e1 := evaluate(pots[l-3:], rules, 4)
	e2 := evaluate(pots[l-4:], rules, 4)

	s1 := evaluate(pots[:3], rules, 2)
	s2 := evaluate(pots[:4], rules, 3)

	tmp := [3]int{0, s1, s2}
	tmpI := 0

	// evaluate rest inbetween
	var x int
	for x = 2; x < l-2; x++ {
		tmp[tmpI] = evaluate(pots[x-2:x+3], rules, 4)
		tmpI = (tmpI + 1) % len(tmp)
		pots[x-2] = tmp[tmpI]
	}

	tmpI++
	pots[x-2] = tmp[tmpI%len(tmp)]
	tmpI++
	pots[x-1] = tmp[tmpI%len(tmp)]
	pots[x] = e2
	pots[x+1] = e1
}
