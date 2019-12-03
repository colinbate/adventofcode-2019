package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var (
		input string
		total int64
		n     int
	)
	n, _ = fmt.Scanln(&input)
	for n > 0 {
		var fuel int64
		mass, _ := strconv.ParseInt(input, 10, 64)
		for mass > 0 {
			mass = int64(math.Max(float64((mass/3)-2), 0))
			fuel += mass
		}
		total += fuel
		n, _ = fmt.Scanln(&input)
	}
	fmt.Println("Total:", total)
}
