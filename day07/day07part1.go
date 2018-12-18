package day07

import (
	"fmt"
	"sort"
)

// part1
func part1(tuples []*instructionTuple, ch chan<- string) {
	// map of children with their corresponding parents
	nodes := initNodes(tuples)

	// result string
	result := make([]rune, len(nodes))

	// checks who is eligible in current round to finish
	// (has every parent done his part?)
	thisRound := []rune{}

	for i := 0; i < len(result); i++ {
	nodeLoop:
		for child, parents := range nodes {
			// was child already added to result string?
			if arrContains(child, result[:i]) || arrContains(child, thisRound) {
				continue nodeLoop
			}

			// have all parents of child node done their part?
			for _, p := range parents {
				if !arrContains(p, result[:i]) {
					continue nodeLoop
				}
			}

			// child is eligible to be added to string! will he be next?
			thisRound = append(thisRound, child)
		}

		// only the first can come through
		sort.Slice(thisRound, func(i, j int) bool {
			return thisRound[i] < thisRound[j]
		})

		result[i] = thisRound[0]
		thisRound = thisRound[1:]
	}

	ch <- fmt.Sprintf("Part 1: %s", string(result))
}

func arrContains(element rune, arr []rune) bool {
	for _, r := range arr {
		if r == element {
			return true
		}
	}

	return false
}

func initNodes(tuples []*instructionTuple) map[rune][]rune {
	nodes := map[rune][]rune{}

	for _, t := range tuples {
		// add parent to map
		if _, ok := nodes[t.first]; !ok {
			nodes[t.first] = []rune{}
		}

		// add child to map
		if _, ok := nodes[t.then]; !ok {
			nodes[t.then] = []rune{}
		}
		// add parent to child
		nodes[t.then] = append(nodes[t.then], t.first)
	}

	return nodes
}
