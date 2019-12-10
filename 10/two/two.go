package two

import (
	"fmt"
	"math"
	"sort"
)

type asteroid struct {
	x, y      int
	dist, ang float64
	zapped    bool
}

type asteroidField []asteroid

const newline = 10
const asteroidChar = 35

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (to *asteroid) setDist(fromx int, fromy int) {
	dx := to.x - fromx
	dy := to.y - fromy
	to.dist = math.Sqrt(float64(dx*dx + dy*dy))
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

func slopeToDeg(slope [2]int) float64 {
	deg := math.Atan2(float64(slope[0]), float64(-slope[1])) * 180.0 / math.Pi
	if deg < 0 {
		deg += 360
	}
	return deg
}

func (to *asteroid) setAngle(fromX int, fromY int) {
	dx := to.x - fromX
	dy := to.y - fromY
	slope := reduceFrac(dx, dy)
	deg := slopeToDeg(slope)
	to.ang = deg
}

func (f asteroidField) Len() int {
	return len(f)
}

func (f asteroidField) Less(i, j int) bool {
	a := (f)[i]
	b := (f)[j]
	if a.ang == b.ang {
		return a.dist < b.dist
	}
	return a.ang < b.ang
}

func (f asteroidField) Swap(i, j int) {
	a := (f)[i]
	(f)[i] = (f)[j]
	(f)[j] = a
}

// These are the coords of the winner from Part One
const originX = 23
const originY = 29

// Run is the entry point for this solution.
func Run() {
	var (
		field     asteroidField
		char      int32
		iy, ix    int
		zapCount  int
		iter      int
		acount    int
		found     asteroid
		lastAngle float64 = 1000
	)
	fmt.Println("Part Two")
	field = make(asteroidField, 0, 35*35)
	for {
		n, _ := fmt.Scanf("%c", &char)
		if n == 0 {
			break
		}
		if char == newline {
			iy++
			ix = 0
		} else if char == asteroidChar && !(ix == originX && iy == originY) {
			newa := asteroid{ix, iy, 0, 0, false}
			newa.setDist(originX, originY)
			newa.setAngle(originX, originY)
			field = append(field, newa)
		}
		if char != newline {
			ix++
		}
	}
	iy++
	sort.Sort(field)
	acount = len(field)
	for zapCount < acount {
		cind := iter % acount
		ast := field[cind]
		if !ast.zapped && ast.ang != lastAngle {
			ast.zapped = true
			zapCount++
			lastAngle = ast.ang
			if zapCount == 200 {
				found = ast
			}
		}
		iter++
	}
	fmt.Println("200th", found.x*100+found.y)
}
