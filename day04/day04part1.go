package day04

import (
	"fmt"
)

// part1
func part1(logs []*loggedAction, ch chan<- string) {

	sleepiestGuard := getLongestSleepingGuard(logs)
	sleepiestMinute := getSleepiestMinute(logs, sleepiestGuard)

	ch <- fmt.Sprintf("Part 1: Guard #%d has been mostly asleep at minute %d. Multiplied it gives you %d",
		sleepiestGuard, sleepiestMinute, sleepiestGuard*sleepiestMinute)
}

// getLongestSleepingGuard returns guard ID who has been asleep the longest
func getLongestSleepingGuard(logs []*loggedAction) int {
	timeSlept := map[int]int{}

	lastAction := wakesUp
	var currentGuard int
	var timeFallenAsleep int
	for _, logged := range logs {

		// i'm barely checking anything, just hoping that I'm getting valid input
		if logged.action == beginsShift && (lastAction == wakesUp || lastAction == beginsShift) {
			currentGuard = logged.guard

		} else if logged.action == fallsAsleep && (lastAction == wakesUp || lastAction == beginsShift) {
			timeFallenAsleep = logged.minute

		} else if logged.action == wakesUp && lastAction == fallsAsleep {
			timeSlept[currentGuard] += logged.minute - timeFallenAsleep

		}
		lastAction = logged.action
	}

	maxGuard := 0
	maxSlept := 0
	for guardID, slept := range timeSlept {
		if slept > maxSlept {
			maxGuard = guardID
			maxSlept = slept
		}
	}

	return maxGuard
}

// getSleepiestMinute counts from minute index 0-59 how often sleepiestGuard has been asleep
func getSleepiestMinute(logs []*loggedAction, sleepiestGuard int) int {
	minutes := [60]int{}
	maxMinutesIndex := 0
	maxMinutes := 0

	curGuard := -1
	var timeFallenAsleep int
	for _, l := range logs {
		if curGuard != sleepiestGuard {
			if l.action == beginsShift {
				curGuard = l.guard
			}

			continue
		}

		if l.action == beginsShift {
			curGuard = l.guard

		} else if l.action == fallsAsleep {
			timeFallenAsleep = l.minute

		} else {
			for min := timeFallenAsleep; min < l.minute; min++ {
				minutes[min]++
				if minutes[min] > maxMinutes {
					maxMinutesIndex = min
					maxMinutes = minutes[min]
				}
			}
		}
	}

	return maxMinutesIndex
}
