package day08

import "fmt"

// part2
func part2(root *node, ch chan<- string) {

	sum := calcExtendedSum(root)

	ch <- fmt.Sprintf("Part 2: %d", sum)
}

// almost the same as calcSum with the added rule that
// those nodes without children will count their metadata
// and those nodes with children will count those children at metadata[x].
// if x is out of bounds, it doesn't count
func calcExtendedSum(n *node) int {
	sum := 0

	if len(n.children) == 0 {
		for _, meta := range n.metadata {
			sum += meta
		}

	} else {
		for _, meta := range n.metadata {
			if meta <= len(n.children) {
				sum += calcExtendedSum(n.children[meta-1])
			}
		}
	}

	return sum
}
