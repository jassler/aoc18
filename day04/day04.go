package day04

import (
	"fmt"
	"log"
	"sort"

	"github.com/jassler/aoc18/fileparser"
)

type actionType int

const (
	beginsShift actionType = iota
	fallsAsleep
	wakesUp
)

type loggedAction struct {
	action actionType
	minute int
	guard  int
}

var logs []*loggedAction

// Start explores this problem set: https://adventofcode.com/2018/day/4
func Start(inputPath string, ch chan<- string) {

	lines := fileparser.ToStringArray(inputPath)
	sort.Strings(lines)

	logs = make([]*loggedAction, len(lines))
	for x := 0; x < len(lines); x++ {
		err := parseDate(lines[x], x)
		if err != nil {
			log.Fatalf("Couldn't parse line \"%s\" in line %d: %v", lines[x], x, err)
		}
	}

	go part1(logs, ch)
	go part2(logs, ch)
}

func parseDate(line string, index int) error {
	// [1518-08-23 00:39] wakes up
	var minute int
	var action actionType

	minuteIndex := len("[1518-08-23 00:")
	actionIndex := len("[1518-11-12 00:00] ")

	_, err := fmt.Sscanf(line[minuteIndex:], "%d", &minute)
	if err != nil {
		return fmt.Errorf("\"%s\" produced the following error while being parsed: %v", line, err)
	}

	curGuard := -1
	changeGuard := func(newGuard int) {
		curGuard = newGuard
	}

	switch line[actionIndex:] {
	case "falls asleep":
		action = fallsAsleep
	case "wakes up":
		action = wakesUp
	default:

		action = beginsShift
		var newGuard int
		_, err = fmt.Sscanf(line[actionIndex:], "Guard #%d", &newGuard)
		if err != nil {
			return err
		}
		changeGuard(newGuard)
	}
	logs[index] = &loggedAction{
		action: action,
		minute: minute,
		guard:  curGuard,
	}

	return nil
}
