package day08

import (
	"strconv"
	"strings"

	"github.com/jassler/aoc18/fileparser"
)

type node struct {
	children []*node
	metadata []int
}

// Start explores this problem set: https://adventofcode.com/2018/day/8
func Start(inputPath string, ch chan<- string) {
	numStr := strings.Split(fileparser.ToStringArray(inputPath)[0], " ")
	nums := make([]int, len(numStr))

	for i := 0; i < len(nums); i++ {
		nums[i], _ = strconv.Atoi(numStr[i])
	}

	// main challenge, generating the tree
	root := generateTree(nums)

	go part1(root, ch)
	go part2(root, ch)
}

func generateTree(nums []int) *node {
	// luckily it already starts with the root element
	root, _ := generateChild(nums)
	return root
}

func generateChild(nums []int) (*node, int) {

	// metadata will be added at the end
	// length of metadata is stored in nums[1]
	n := &node{
		children: make([]*node, nums[0]),
	}

	// index 0 and 1 contain length of children / metadata
	// each child should start at the second index
	i := 2
	var newI int

	for x := 0; x < len(n.children); x++ {

		// pass slice that starts at index i so they can start counting with index 0
		n.children[x], newI = generateChild(nums[i : len(nums)-nums[0]-1])
		i += newI
	}

	// add metadata slice
	lastIndex := i + nums[1]
	n.metadata = nums[i:lastIndex]

	return n, lastIndex
}
