package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Vertex struct {
	X int
	Y int
}

type Segment struct {
	Start Vertex
	End   Vertex
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isVertical(seg Segment) bool {
	return seg.Start.X == seg.End.X
}

func between(target int, end1 int, end2 int) bool {
	return (end1 <= target && target <= end2) || (end2 <= target && target <= end1)
}

func overlap(s1 int, e1 int, s2 int, e2 int) (bool, int) {
	var best int = math.MaxInt32
	if between(s1, s2, e2) && s1 < best {
		best = s1
	}
	if between(e1, s2, e2) && e1 < best {
		best = e1
	}
	if between(s2, s1, e1) && s2 < best {
		best = s2
	}
	if between(e2, s1, e1) && e2 < best {
		best = e2
	}
	found := best != math.MaxInt32
	return found, best
}

func intersects(s1 Segment, s2 Segment) (bool, Vertex) {
	var (
		vert Segment
		horz Segment
	)
	v1 := isVertical(s1)
	v2 := isVertical(s2)
	if v1 == v2 {
		if v1 && s1.Start.X == s2.Start.X {
			hit, best := overlap(s1.Start.Y, s1.End.Y, s2.Start.Y, s2.End.Y)
			if hit {
				return true, Vertex{s1.Start.X, best}
			}
		}
		if !v1 && s1.Start.Y == s2.Start.Y {
			hit, best := overlap(s1.Start.X, s1.End.X, s2.Start.X, s2.End.X)
			if hit {
				return true, Vertex{best, s1.Start.Y}
			}
		}
		return false, Vertex{}
	}
	if v1 {
		vert = s1
		horz = s2
	} else {
		vert = s2
		horz = s1
	}
	if between(vert.Start.X, horz.Start.X, horz.End.X) && between(horz.Start.Y, vert.Start.Y, vert.End.Y) {
		return true, Vertex{vert.Start.X, horz.Start.Y}
	}
	return false, Vertex{}
}

func distance(p Vertex) int {
	return abs(p.X) + abs(p.Y)
}

func findEnd(start Vertex, delta string) Vertex {
	var (
		dir   string
		value int
		end   Vertex
	)
	dir = string(delta[0])
	value, _ = strconv.Atoi(strings.TrimLeft(delta, "RLUD"))
	switch dir {
	case "R":
		end = Vertex{start.X + value, start.Y}
	case "L":
		end = Vertex{start.X - value, start.Y}
	case "D":
		end = Vertex{start.X, start.Y - value}
	case "U":
		end = Vertex{start.X, start.Y + value}
	}
	return end
}

func main() {
	var (
		first   string
		second  string
		deltasA []string
		deltasB []string
		segsA   []Segment
		segsB   []Segment
		current Vertex
		closest int
	)
	fmt.Scanln(&first)
	fmt.Scanln(&second)

	deltasA = strings.Split(first, ",")
	deltasB = strings.Split(second, ",")
	segsA = make([]Segment, 0, len(deltasA)+1)
	segsB = make([]Segment, 0, len(deltasB)+1)

	fmt.Printf("Found %d of A and %d of B\n", len(deltasA), len(deltasB))

	current = Vertex{0, 0}
	for _, d := range deltasA {
		endPoint := findEnd(current, d)
		seg := Segment{current, endPoint}
		segsA = append(segsA, seg)
		current = endPoint
		fmt.Printf("Current A: (%d, %d)\n", current.X, current.Y)
	}

	current = Vertex{0, 0}
	for _, d := range deltasB {
		endPoint := findEnd(current, d)
		seg := Segment{current, endPoint}
		segsB = append(segsB, seg)
		current = endPoint
		fmt.Printf("Current B: (%d, %d)\n", current.X, current.Y)

	}

	closest = math.MaxInt32
	for i, s := range segsA {
		for j, ss := range segsB {
			cross, where := intersects(s, ss)
			if cross && !(i == 0 && j == 0 && where.X == 0 && where.Y == 0) {
				fmt.Printf("Found intersection at (%d, %d)\n", where.X, where.Y)
				dist := distance(where)
				if dist < closest {
					closest = dist
				}
			}
		}
	}
	fmt.Println("Closest:", closest)
}
