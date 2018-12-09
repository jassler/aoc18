package day04

import (
	"fmt"
)

// part2
func part2(logs []*loggedAction, ch chan<- string) {

	guardSleepCycle := map[int][60]int{}

	var timeFallenAsleep int

	maxGuardID := 0
	maxMinuteCount := 0
	maxMinute := 0

	for _, l := range logs {
		if l.action == fallsAsleep {
			timeFallenAsleep = l.minute

		} else if l.action == wakesUp {
			for min := timeFallenAsleep; min < l.minute; min++ {
				cycle, ok := guardSleepCycle[l.guard]
				if ok {
					cycle[min]++
				} else {
					cycle = [60]int{}
					cycle[min] = 1
				}
				guardSleepCycle[l.guard] = cycle
				if cycle[min] > maxMinuteCount {
					maxMinuteCount = cycle[min]
					maxMinute = min
					maxGuardID = l.guard
				}
			}
		}
	}

	ch <- fmt.Sprintf("Part 2: Guard #%d has been most asleep at minute %d. %d * %d = %d",
		maxGuardID, maxMinute, maxGuardID, maxMinute, maxGuardID*maxMinute)
}
