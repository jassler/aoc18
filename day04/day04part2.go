package day04

import (
	"fmt"
	"time"
)

// part2
func part2(logs []*loggedAction, ch chan<- string) {

	guardSleepCycle := map[int][60]int{}

	var timeFallenAsleep *time.Time
	curGuard := -1

	maxGuardID := 0
	maxMinuteCount := 0
	maxMinute := 0

	for _, l := range logs {
		if l.action == beginsShift {
			curGuard = l.guard

		} else if l.action == fallsAsleep {
			timeFallenAsleep = &l.timestamp

		} else {
			for min := timeFallenAsleep.Minute(); min < l.timestamp.Minute(); min++ {
				cycle, ok := guardSleepCycle[curGuard]
				if ok {
					cycle[min]++
				} else {
					cycle = [60]int{}
					cycle[min] = 1
				}
				guardSleepCycle[curGuard] = cycle
				if cycle[min] > maxMinuteCount {
					maxMinuteCount = cycle[min]
					maxMinute = min
					maxGuardID = curGuard
				}
			}
		}
	}

	ch <- fmt.Sprintf("Part 2: Guard #%d has been most asleep at minute %d. %d * %d = %d",
		maxGuardID, maxMinute, maxGuardID, maxMinute, maxGuardID*maxMinute)
}
