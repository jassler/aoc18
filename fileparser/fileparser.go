package fileparser

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// ToStringArray stores each line from a file in a string array
func ToStringArray(filepath string) []string {
	strings := make([]string, 0)

	toArray(filepath, func(line string) error {
		strings = append(strings, line)
		return nil
	})

	return strings
}

// ToIntArray converts each line of a file into an integer and stores it in an array
func ToIntArray(filepath string) []int {
	numbers := make([]int, 0)

	toArray(filepath, func(line string) error {
		num, err := strconv.Atoi(line)
		if err != nil {
			return errors.New(line + " could not be converted to a number!")
		}
		numbers = append(numbers, num)
		return nil
	})

	return numbers
}

// toArray opens a file and calls function f on each line
// if an error appears the program shuts down (panic)
func toArray(filepath string, f func(string) error) {

	// open file
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0

	// go through each line and call function f
	for scanner.Scan() {
		lineNum++

		err := f(scanner.Text())
		if err != nil {
			panic(fmt.Sprintf("Error in line %d: %s", lineNum, err.Error()))
		}
	}
}
