package day04

import (
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/jassler/aoc18/fileparser"
)

type actionType int

const (
	beginsShift actionType = iota
	fallsAsleep
	wakesUp
)

type loggedAction struct {
	timestamp time.Time
	action    actionType
	guard     int
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
	// cheating a little bit. I'm ignoring leap years and other stuff
	year := 2016
	var month, day, hour, minute int
	count, err := fmt.Sscanf(line, "[1518-%d-%d %d:%d", &month, &day, &hour, &minute)
	if err != nil {
		return fmt.Errorf("\"%s\" produced the following error while being parsed: %v", line, err)
	}
	if count != 4 {
		return fmt.Errorf("\"%s\" doesn't have 4 numbers to parse, instead got %d", line, count)
	}

	a := loggedAction{
		timestamp: time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC),
		guard:     -1,
	}

	switch line[19:] {
	case "falls asleep":
		a.action = fallsAsleep
	case "wakes up":
		a.action = wakesUp
	default:

		a.action = beginsShift
		_, err = fmt.Sscanf(line[26:], "%d", &a.guard)
		if err != nil {
			return err
		}
	}
	logs[index] = &a

	return nil
}
