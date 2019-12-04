package main

import (
	"fmt"
	"math"
)

func ipow(n int) int {
	return int(math.Pow10(n))
}

func getDigit(num int, digit int) int {
	// From right to left 0-based
	if digit == 0 {
		return num % 10
	}
	return (num / ipow(digit)) % 10
}

func hasDouble(num int) bool {
	var val int = getDigit(num, 5)
	for i := 1; i < 6; i++ {
		curr := getDigit(num, 5-i)
		if curr == val {
			return true
		}
		val = curr
	}
	return false
}

func increment(num int) int {
	if num < 9 {
		return num + 1
	}
	if num == 9 {
		return num // Stop at all nines
	}
	ones := getDigit(num, 0)
	if ones != 9 {
		return num + 1
	}
	higher := num / 10
	higher = increment(higher)
	return higher*10 + (higher % 10)
}

const stop int = 732736

func main() {
	var (
		odo   int = 256666
		count int
	)
	for odo < stop {
		if hasDouble(odo) {
			count++
		}
		odo = increment(odo)
	}

	fmt.Println("Count:", count)
}
