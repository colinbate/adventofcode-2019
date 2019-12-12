package one

import (
	"fmt"
)

type moon struct {
	name                string
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

func (m *moon) applyGravity(moons system) {
	for _, other := range moons {
		if other.name != (*m).name {
			(*m).vx += grav((*m).x, other.x)
			(*m).vy += grav((*m).y, other.y)
			(*m).vz += grav((*m).z, other.z)
		}
	}
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

// Run is the entry point for this solution.
func Run() {
	var (
		moons  system
		iter   int
		energy int
	)
	fmt.Println("Part One")
	// Input, too lazy to read it in
	moons = system{
		moon{"A", 14, 9, 14, 0, 0, 0},
		moon{"B", 9, 11, 6, 0, 0, 0},
		moon{"C", -6, 14, -4, 0, 0, 0},
		moon{"D", 4, -4, -3, 0, 0, 0},
	}
	mlen := len(moons)
	for iter < 1000 {
		for mi := 0; mi < mlen; mi++ {
			moons[mi].applyGravity(moons)
		}
		for mi := 0; mi < mlen; mi++ {
			moons[mi].applyVelocity()
		}
		iter++
	}
	energy = moons.getEnergy()
	fmt.Println("Total energy:", energy)
}
