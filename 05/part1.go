package main

import (
	"fmt"
	"strconv"
	"strings"
)

func toCodes(vals []string) []int {
	codes := make([]int, len(vals))
	for i, s := range vals {
		codes[i], _ = strconv.Atoi(s)
	}
	return codes
}

func getOpCode(code int) (int, [3]int) {
	var modes [3]int
	op := code % 100
	modes[0] = code / 100 % 10
	modes[1] = code / 1000 % 10
	modes[2] = code / 10000 % 10
	return op, modes
}

const ADD int = 1
const MULT int = 2
const INPUT int = 3
const OUTPUT int = 4
const HALT int = 99

func computer(codes []int, input []int) {
	var (
		inputPtr int
		ptr      int
		opcode   int
		modes    [3]int
	)

	for ptr < len(codes) {
		opcode, modes = getOpCode(codes[ptr])
		switch opcode {
		case ADD:
			arg1 := codes[ptr+1]
			arg2 := codes[ptr+2]
			if modes[0] == 0 {
				arg1 = codes[arg1]
			}
			if modes[1] == 0 {
				arg2 = codes[arg2]
			}
			out := codes[ptr+3]
			codes[out] = arg1 + arg2
			ptr += 4
		case MULT:
			arg1 := codes[ptr+1]
			arg2 := codes[ptr+2]
			if modes[0] == 0 {
				arg1 = codes[arg1]
			}
			if modes[1] == 0 {
				arg2 = codes[arg2]
			}
			out := codes[ptr+3]
			codes[out] = arg1 * arg2
			ptr += 4
		case INPUT:
			arg := codes[ptr+1]
			codes[arg] = input[inputPtr]
			inputPtr++
			ptr += 2
		case OUTPUT:
			arg := codes[ptr+1]
			if modes[0] == 0 {
				arg = codes[arg]
			}
			fmt.Println(arg)
			ptr += 2
		case HALT:
			fmt.Println("halt")
			return
		}
	}
}

func main() {
	var (
		input string
		codes []int
	)
	fmt.Scanln(&input)
	codes = toCodes(strings.Split(input, ","))

	computer(codes, []int{1})
}
