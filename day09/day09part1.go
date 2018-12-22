package day09

import "fmt"

type marble struct {
	points     int
	next, prev *marble
}

// part1
func part1(players, lastMarble int, ch chan<- string) {

	scores := simulateGame(players, lastMarble)
	player, score := findWinningPlayer(scores)

	ch <- fmt.Sprintf("Part 1: Player %d won with %d points", player, score)
}

// simulateGame with the strange circle rules. Returns resulting score of each player
func simulateGame(players, lastMarble int) []int {
	curPlayer := 0
	scores := make([]int, players)

	curMarble := &marble{
		points: 0,
	}
	curMarble.next = curMarble
	curMarble.prev = curMarble

	for points := 1; points <= lastMarble; points++ {

		if points%23 == 0 {
			// elf gets points!
			// go seven to the left
			for i := 0; i < 7; i++ {
				curMarble = curMarble.prev
			}
			scores[curPlayer] += points + curMarble.points

			// remove current marble
			curMarble.prev.next = curMarble.next
			curMarble.next.prev = curMarble.prev
			curMarble = curMarble.next
		} else {
			// elf adds marble to circle
			curMarble = curMarble.next

			nextMarble := &marble{
				points: points,
				next:   curMarble.next,
				prev:   curMarble,
			}

			curMarble.next.prev = nextMarble
			curMarble.next = nextMarble
			curMarble = nextMarble
		}

		// next player's turn!
		// Modulo operation makes it loop back to player 0 when last player took his turn
		curPlayer = (curPlayer + 1) % players
	}

	return scores
}

// findWinningPlayer returns player number with corresponding points
// player number is already incremented by 1
func findWinningPlayer(scores []int) (int, int) {
	var index, score int
	for i, s := range scores {
		if s > score {
			index = i
			score = s
		}
	}

	return index + 1, score
}
