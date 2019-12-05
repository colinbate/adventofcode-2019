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

func getArg(codes []int, ptr int, mode int) int {
	arg := codes[ptr]
	if mode == 0 {
		arg = codes[arg]
	}
	return arg
}

const ADD int = 1
const MULT int = 2
const INPUT int = 3
const OUTPUT int = 4
const JUMP_IF_TRUE int = 5
const JUMP_IF_FALSE int = 6
const LESS_THAN int = 7
const EQUALS int = 8
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
			arg1 := getArg(codes, ptr+1, modes[0])
			arg2 := getArg(codes, ptr+2, modes[1])
			out := codes[ptr+3]
			codes[out] = arg1 + arg2
			ptr += 4
		case MULT:
			arg1 := getArg(codes, ptr+1, modes[0])
			arg2 := getArg(codes, ptr+2, modes[1])
			out := codes[ptr+3]
			codes[out] = arg1 * arg2
			ptr += 4
		case INPUT:
			arg := codes[ptr+1]
			codes[arg] = input[inputPtr]
			inputPtr++
			ptr += 2
		case OUTPUT:
			arg := getArg(codes, ptr+1, modes[0])
			fmt.Println(arg)
			ptr += 2
		case JUMP_IF_TRUE:
			arg1 := getArg(codes, ptr+1, modes[0])
			arg2 := getArg(codes, ptr+2, modes[1])
			if arg1 != 0 {
				ptr = arg2
			} else {
				ptr += 3
			}
		case JUMP_IF_FALSE:
			arg1 := getArg(codes, ptr+1, modes[0])
			arg2 := getArg(codes, ptr+2, modes[1])
			if arg1 == 0 {
				ptr = arg2
			} else {
				ptr += 3
			}
		case LESS_THAN:
			arg1 := getArg(codes, ptr+1, modes[0])
			arg2 := getArg(codes, ptr+2, modes[1])
			out := codes[ptr+3]
			if arg1 < arg2 {
				codes[out] = 1
			} else {
				codes[out] = 0
			}
			ptr += 4
		case EQUALS:
			arg1 := getArg(codes, ptr+1, modes[0])
			arg2 := getArg(codes, ptr+2, modes[1])
			out := codes[ptr+3]
			if arg1 == arg2 {
				codes[out] = 1
			} else {
				codes[out] = 0
			}
			ptr += 4
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

	computer(codes, []int{5})
}
