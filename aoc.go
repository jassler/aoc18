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
	"github.com/jassler/aoc18/day09"
)

var functions = map[string]func(input string, ch chan<- string){
	"01": day01.Start,
	"02": day02.Start,
	"03": day03.Start,
	"04": day04.Start,
	"05": day05.Start,
	"06": day06.Start,
	"07": day07.Start,
	"08": day08.Start,
	"09": day09.Start,
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
	if len(day) == 1 {
		day = "0" + day
	}

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
