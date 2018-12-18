package day07

import (
	"fmt"
	"sort"
)

type worker struct {
	letter rune
	second int
}

// part2
func part2(steps []*instructionTuple, ch chan<- string) {
	// part 1 function to generate map of children with their parents
	nodes := initNodes(steps)

	// result string
	result := make([]rune, len(nodes))
	// current letter we're filling in
	resultIndex := 0

	// current rune workers are working with
	workers := make([]*worker, 5)
	for i := range workers {
		workers[i] = &worker{
			letter: 0,
			second: 0,
		}
	}

	// first step is at second 0
	// -> after first increment, seconds has to be 0
	seconds := -1

	for resultIndex < len(result) {
		// one second passes
		done := simulateStep(workers, nodes, result[:resultIndex])
		seconds++

		if len(done) == 0 {
			continue
		}

		// add letters that were done to our result text
		sort.Slice(done, func(i, j int) bool {
			return done[i] < done[j]
		})

		for _, r := range done {
			result[resultIndex] = r
			resultIndex++
		}

	}

	ch <- fmt.Sprintf("Part 2: It takes %d seconds (new result is %s)", seconds, string(result))
}

func simulateStep(workers []*worker, nodes map[rune][]rune, done []rune) []rune {
	// to calculate seconds
	// eg. if letter 'A' appears, it will take 'A'+diff = 61 seconds
	const diff = 60 - 'A' + 1

	// workers finished in this round with the following letters...
	thisRound := []rune{}

	// workers who can start a new job
	// either from first time or when worker is done working with his letter
	idle := []*worker{}
	for _, work := range workers {
		// is worker idle?
		if work.letter == 0 {
			idle = append(idle, work)
		} else {
			work.second--
			if work.second == 0 {
				thisRound = append(thisRound, work.letter)
				idle = append(idle, work)
				work.letter = 0
			}
		}
	}

	if len(idle) > 0 {
		idleIndex := 0

	nodeLoop:
		for child, parents := range nodes {
			// was child already added to result string?
			if arrContains(child, done) || arrContains(child, thisRound) {
				continue nodeLoop
			}

			// is letter already being worked on by another worker
			for _, w := range workers {
				if w.letter == child {
					continue nodeLoop
				}
			}

			// have all parents of child node done their part?
			for _, p := range parents {
				if !(arrContains(p, done) || arrContains(p, thisRound)) {
					continue nodeLoop
				}
			}

			// new work for idle worker!
			idle[idleIndex].letter = child
			idle[idleIndex].second = int(child) + diff
			idleIndex++

			// do all workers have something to do?
			if idleIndex >= len(idle) {
				break nodeLoop
			}
		}
	}

	return thisRound
}
