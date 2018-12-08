package day01

import "fmt"

// basic premis is: calculate sum of each row
func part1(numbers []int, ch chan<- string) {
	result := 0

	for _, num := range numbers {
		result += num
	}

	ch <- fmt.Sprint("Part 1: ", result)
}
