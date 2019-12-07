package two

import (
	"fmt"
	"strconv"
	"strings"
)

// Program is a sequence of instructinos for intcode computer
type Program []int

// Copy returns a copy of the Program.
func (s Program) Copy() Program {
	copy := make(Program, 0, len(s))
	return append(copy, s...)
}

func toCodes(vals []string) Program {
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

// ADD adds
const ADD int = 1

// MULT multiplies
const MULT int = 2

// INPUT inputs
const INPUT int = 3

// OUTPUT outputs
const OUTPUT int = 4

// JUMPIFTRUE does that
const JUMPIFTRUE int = 5

// JUMPIFFALSE does that
const JUMPIFFALSE int = 6

// LESSTHAN compares inequality
const LESSTHAN int = 7

// EQUALS compares equality
const EQUALS int = 8

// HALT halts
const HALT int = 99

func computer(name string, codes []int, input chan int, output chan int, out2 chan int) {
	var (
		ptr    int
		opcode int
		modes  [3]int
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
			// fmt.Println("computer", name, "wants input")
			val, ok := <-input
			if ok {
				codes[arg] = val
				//fmt.Println("computer", name, "received input", val)
				//} else {
				//	fmt.Println("computer", name, "input not ok")
			}
			ptr += 2
		case OUTPUT:
			arg := getArg(codes, ptr+1, modes[0])
			//fmt.Println("computer", name, "outputting", arg)
			output <- arg
			if out2 != nil {
				out2 <- arg
			}
			ptr += 2
		case JUMPIFTRUE:
			arg1 := getArg(codes, ptr+1, modes[0])
			arg2 := getArg(codes, ptr+2, modes[1])
			if arg1 != 0 {
				ptr = arg2
			} else {
				ptr += 3
			}
		case JUMPIFFALSE:
			arg1 := getArg(codes, ptr+1, modes[0])
			arg2 := getArg(codes, ptr+2, modes[1])
			if arg1 == 0 {
				ptr = arg2
			} else {
				ptr += 3
			}
		case LESSTHAN:
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
			// fmt.Println("halt", name)
			close(output)
			if out2 != nil {
				close(out2)
			}
			return
		}
	}
	return
}

func ampChain(codes Program, phases []int) int {
	//fmt.Println("Starting ampChain for", phases)
	var result int
	ab := make(chan int)
	bc := make(chan int)
	cd := make(chan int)
	de := make(chan int)
	ea := make(chan int, 1)
	eout := make(chan int)
	go computer("A", codes.Copy(), ea, ab, nil)
	go computer("B", codes.Copy(), ab, bc, nil)
	go computer("C", codes.Copy(), bc, cd, nil)
	go computer("D", codes.Copy(), cd, de, nil)
	go computer("E", codes.Copy(), de, ea, eout)
	ea <- phases[0]
	ab <- phases[1]
	bc <- phases[2]
	cd <- phases[3]
	de <- phases[4]
	ea <- 0 // Initial value
	for out := range eout {
		result = out
	}
	return result
}

// Thank you Heaps algorithm
func generatePermutations(list []int, size int, codes Program, results chan int) {
	if size == 1 {
		//fmt.Println(list)
		results <- ampChain(codes, list)
	}
	for i := 0; i < size; i++ {
		generatePermutations(list, size-1, codes, results)

		if size%2 == 1 {
			temp := list[0]
			list[0] = list[size-1]
			list[size-1] = temp
		} else {
			temp := list[i]
			list[i] = list[size-1]
			list[size-1] = temp
		}
	}
}

func evaluate(codes Program, results chan int) {
	generatePermutations([]int{5, 6, 7, 8, 9}, 5, codes, results)
	close(results)
}

// Run is the entry point for this solution.
func Run() {
	fmt.Println("Part One")
	var (
		input   string
		codes   Program
		largest int
		results chan int
	)
	fmt.Scanln(&input)
	codes = toCodes(strings.Split(input, ","))
	results = make(chan int)
	go evaluate(codes, results)
	for r := range results {
		if r > largest {
			largest = r
		}
	}
	//largest = ampChain(codes, []int{9, 8, 7, 6, 5})
	fmt.Println("Largest:", largest)
}
