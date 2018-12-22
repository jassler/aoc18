package day09

import "fmt"

// part2
func part2(players, lastMarble int, ch chan<- string) {

	scores := simulateGame(players, lastMarble*100)
	player, score := findWinningPlayer(scores)

	ch <- fmt.Sprintf("Part 2: Player %d won with %d points", player, score)
}
