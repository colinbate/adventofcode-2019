package one

import (
	"fmt"
)

type asteroid struct {
	x, y int
}

const newline = 10
const asteroidChar = 35

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func reduceFrac(dx, dy int) [2]int {
	div := gcd(abs(dy), abs(dx))
	dx = dx / div
	dy = dy / div
	return [2]int{dx, dy}
}

func findDetectable(field []asteroid, fromX int, fromY int) int {
	var (
		count         int
		blockedSlopes map[[2]int]bool
	)
	blockedSlopes = make(map[[2]int]bool)
	for _, roid := range field {
		if roid.x == fromX && roid.y == fromY {
			continue
		}
		dx := roid.x - fromX
		dy := roid.y - fromY
		slope := reduceFrac(dx, dy)
		if !blockedSlopes[slope] {
			blockedSlopes[slope] = true
			count++
		}
	}
	return count
}

// Run is the entry point for this solution.
func Run() {
	var (
		field  []asteroid
		char   int32
		iy, ix int
		most   int
		best   asteroid
	)
	fmt.Println("Part One")
	field = make([]asteroid, 0, 35*35)
	for {
		n, _ := fmt.Scanf("%c", &char)
		if n == 0 {
			break
		}
		if char == newline {
			iy++
			ix = 0
		} else if char == asteroidChar {
			field = append(field, asteroid{ix, iy})
		}
		if char != newline {
			ix++
		}
	}
	iy++
	for _, r := range field {
		detect := findDetectable(field, r.x, r.y)
		if detect > most {
			most = detect
			best = r
		}
	}
	fmt.Println("Most", most, best)
}
