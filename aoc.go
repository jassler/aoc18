package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/jassler/aoc18/day01"
	"github.com/jassler/aoc18/day02"
	"github.com/jassler/aoc18/day03"
	"github.com/jassler/aoc18/day04"
	"github.com/jassler/aoc18/day05"
	"github.com/jassler/aoc18/day06"
	"github.com/jassler/aoc18/day07"
	"github.com/jassler/aoc18/day08"
)

var functions = map[string]func(input string, ch chan<- string){
	"1": day01.Start,
	"2": day02.Start,
	"3": day03.Start,
	"4": day04.Start,
	"5": day05.Start,
	"6": day06.Start,
	"7": day07.Start,
	"8": day08.Start,
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Specify day number to solve")
		fmt.Println("Currently available:")
		for day := range functions {
			fmt.Println("   ", os.Args[0], day)
		}
		return
	}

	day := os.Args[1]

	inputPath, err := filepath.Abs("input")
	if err != nil {
		panic(err)
	}

	inputFile := path.Join(inputPath, "day"+day+"_input.txt")

	ch := make(chan string, 2)

	functions[day](inputFile, ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
