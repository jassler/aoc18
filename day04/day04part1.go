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

	var timeFallenAsleep int
	for _, logged := range logs {

		// i'm barely checking anything, just hoping that I'm getting valid input
		if logged.action == fallsAsleep {
			timeFallenAsleep = logged.minute

		} else if logged.action == wakesUp {
			timeSlept[logged.guard] += logged.minute - timeFallenAsleep
		}
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

	var timeFallenAsleep int
	for _, l := range logs {
		if l.guard != sleepiestGuard || l.action == beginsShift {
			continue

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
