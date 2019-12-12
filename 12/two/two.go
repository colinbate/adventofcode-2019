package two

import (
	"fmt"
)

type axis struct {
	p, v int
}

type asystem [4]axis

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

func findPeriod(name string, sys asystem, result chan int) {
	iter := 0
	orig := sys.copy()
	for {
		for mi := 0; mi < 4; mi++ {
			for i, other := range sys {
				if mi != i {
					sys[mi].v += grav(sys[mi].p, other.p)
				}
			}
		}
		for mi := 0; mi < 4; mi++ {
			sys[mi].applyVelocity()
		}
		iter++
		if sys == orig {
			fmt.Println(name, "period", iter)
			result <- iter
			break
		}
	}
	return
}

// Run is the entry point for this solution.
func Run() {
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
	results := make(chan int, 3)
	go findPeriod("X", systemX, results)
	go findPeriod("Y", systemY, results)
	go findPeriod("Z", systemZ, results)
	lcmA := lcm(<-results, <-results)
	total := lcm(lcmA, <-results)
	fmt.Println("Total:", total)
}
