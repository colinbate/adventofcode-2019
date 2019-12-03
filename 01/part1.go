package main

import (
	"fmt"
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
		mass, _ := strconv.ParseInt(input, 10, 64)
		fuel := (mass / 3) - 2
		total += fuel
		n, _ = fmt.Scanln(&input)
	}
	fmt.Println("Total:", total)
}
