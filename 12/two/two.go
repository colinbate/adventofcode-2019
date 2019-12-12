package two

import (
	"fmt"
)

type axis struct {
	p, v int
}

type asystem [4]axis

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func grav(me int, other int) (dv int) {
	if me < other {
		dv = 1
	} else if me > other {
		dv = -1
	} else {
		dv = 0
	}
	return
}

func (m *axis) applyVelocity() {
	(*m).p += (*m).v
}

func (sys asystem) copy() asystem {
	newsys := asystem{
		axis{sys[0].p, sys[0].v},
		axis{sys[1].p, sys[1].v},
		axis{sys[2].p, sys[2].v},
		axis{sys[3].p, sys[3].v},
	}
	return newsys
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

// Run is the entry point for this solution.
func Run() {
	var (
		iter int
	)
	fmt.Println("Part Two")
	// Input... too lazy to read it in
	systemX := asystem{
		axis{14, 0},
		axis{9, 0},
		axis{-6, 0},
		axis{4, 0},
	}
	systemY := asystem{
		axis{9, 0},
		axis{11, 0},
		axis{14, 0},
		axis{-4, 0},
	}
	systemZ := asystem{
		axis{14, 0},
		axis{6, 0},
		axis{-4, 0},
		axis{-3, 0},
	}

	allSystem := []asystem{systemX, systemY, systemZ}
	orig := []asystem{systemX.copy(), systemY.copy(), systemZ.copy()}
	periods := [3]int{0, 0, 0}
	for ai, a := range allSystem {
		iter = 0
		for {
			for mi := 0; mi < 4; mi++ {
				for i, other := range a {
					if mi != i {
						a[mi].v += grav(a[mi].p, other.p)
					}
				}
			}
			for mi := 0; mi < 4; mi++ {
				a[mi].applyVelocity()
			}
			iter++
			if a == orig[ai] {
				fmt.Println(ai, "period", iter)
				periods[ai] = iter
				break
			}
		}
	}
	lcmA := lcm(periods[0], periods[1])
	lcmB := lcm(lcmA, periods[2])
	fmt.Println("Total", lcmB)
}
