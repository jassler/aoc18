package day08

import "fmt"

// part1
func part1(root *node, ch chan<- string) {

	sum := calcSum(root)

	ch <- fmt.Sprintf("Part 1: %d", sum)
}

// calculate sum of metadata and sum of its children recursively
func calcSum(n *node) int {
	sum := 0

	// sum of metadata
	for _, meta := range n.metadata {
		sum += meta
	}

	// sum of children RECURSION!
	for _, child := range n.children {
		sum += calcSum(child)
	}

	return sum
}
