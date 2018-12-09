package day06

import "fmt"

// part1
func part1(points []point, ch chan<- string) {
	area := mapArea(points)

	areaIsInfinite := map[byte]bool{}
	squareCount := map[byte]int{}
	for y := 0; y < len(area); y++ {
		for x := 0; x < len(area[y]); x++ {
			b := findClosestPoint(area, point{x: x, y: y})
			if b == 0 {
				continue
			}

			// if on the edge, area is infinite
			if x == 0 || y == 0 || x == (len(area[y])-1) || y == (len(area)-1) {
				areaIsInfinite[b] = true
			} else {
				squareCount[b]++
			}
		}
	}

	largest := 0
	for b, count := range squareCount {
		if !areaIsInfinite[b] && count > largest {
			largest = count
		}
	}

	ch <- fmt.Sprintf("Part 1: Largest non-infinite area is %d", largest)
}

func findClosestPoint(area [][]byte, toPoint point) byte {
	if l := area[toPoint.y][toPoint.x]; l > 0 {
		return l
	}

	// variant of breadth first search
	/*
		suppose a given point 3,3

		depth=2
		the following points need to be traversed
		absolute coordinate | (relative coordinate) | x variable less than depth
		1,3 (-2, 0) x=0
		2,2 (-1,-1) x=1
		3,1 ( 0,-2) x=0
		4,2 ( 1,-1) x=1
		5,3 ( 2, 0) x=0
		4,4 ( 1, 1) x=1
		3,5 ( 0, 2) x=0
		2,4 (-1, 1) x=1
		1,3 (-2, 0) --- loop

		starting with x == 0, it checks straight left, up, right, down
		so we need (-2,0), (0,2), (2,0) and (0,-2)
		(-2,0) in that case can be calculated with (-depth,0)

		with x == 1, we check diagonally
		so we need (-1,-1), (1,-1), (1,1) and (-1,1)
		(-1,-1) in that case can be calculated with (-depth+x, -x)
	*/
	depth := 0
	found := byte(0)
	size := point{
		x: len(area[0]),
		y: len(area),
	}
	for found == 0 {
		depth++
		for x := 0; x < depth; x++ {

			// see calculation at top for each point
			for _, p := range []point{
				point{x: toPoint.x + x - depth, y: toPoint.y - x}, // left to up
				point{x: toPoint.x + x, y: toPoint.y - depth + x}, // up to right
				point{x: toPoint.x - x + depth, y: toPoint.y + x}, // right to down
				point{x: toPoint.x - x, y: toPoint.y + depth - x}, // down to left
			} {
				// point is within bounds?
				if (p.x >= 0) && (p.y >= 0) && (p.x < size.x) && (p.y < size.y) {

					// there's a point saved?
					if l := area[p.y][p.x]; l != 0 {

						// is another item the same distance away?
						if found != 0 {
							return 0
						}
						found = l
					}
				}
			}

		}
	}

	return found
}

// mapArea creates a 2d-map with all points located at map[point.y][point.x]
func mapArea(points []point) [][]byte {
	maxX, maxY := getWidthHeight(points)

	area := make([][]byte, maxY+1)
	for y := 0; y < maxY+1; y++ {
		area[y] = make([]byte, maxX+1)
	}

	for i, p := range points {
		area[p.y][p.x] = byte(i) + 1
	}
	return area
}

func getWidthHeight(points []point) (int, int) {
	maxX, maxY := 0, 0

	for _, p := range points {
		if maxX < p.x {
			maxX = p.x
		}
		if maxY < p.y {
			maxY = p.y
		}
	}

	return maxX, maxY
}
