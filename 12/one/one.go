package one

import (
	"fmt"
)

type moon struct {
	x, y, z, vx, vy, vz int
}

type system []moon

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

func (m *moon) applyVelocity() {
	(*m).x += (*m).vx
	(*m).y += (*m).vy
	(*m).z += (*m).vz
}

func (m *moon) getEnergy() int {
	mm := (*m)
	pot := abs(mm.x) + abs(mm.y) + abs(mm.z)
	kin := abs(mm.vx) + abs(mm.vy) + abs(mm.vz)
	return pot * kin
}

func (sys system) getEnergy() (total int) {
	for _, m := range sys {
		part := m.getEnergy()
		//fmt.Println("Energy from", m.name, "=", part)
		total += part
	}
	return
}

func readMoons() (moons system) {
	var (
		x, y, z int
	)
	moons = make(system, 0)
	for {
		n, _ := fmt.Scanf("<x=%d, y=%d, z=%d>\n", &x, &y, &z)
		if n < 3 {
			break
		}
		moons = append(moons, moon{x, y, z, 0, 0, 0})
	}
	return
}

// Run is the entry point for this solution.
func Run() {
	fmt.Println("Part One")
	moons := readMoons()
	mlen := len(moons)
	for iter := 0; iter < 1000; {
		for mi := 0; mi < mlen; mi++ {
			for i, other := range moons {
				if mi != i {
					moons[mi].vx += grav(moons[mi].x, other.x)
					moons[mi].vy += grav(moons[mi].y, other.y)
					moons[mi].vz += grav(moons[mi].z, other.z)
				}
			}
		}
		for mi := 0; mi < mlen; mi++ {
			moons[mi].applyVelocity()
		}
		iter++
	}
	energy := moons.getEnergy()
	fmt.Println("Total energy:", energy)
}
