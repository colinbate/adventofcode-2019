package one

import (
	"fmt"
	"strconv"
	"strings"
)

// Program is a sequence of instructinos for intcode computer
type Program []int

func reallocate(codes Program, newSize int) Program {
	newProg := make(Program, newSize)
	// The copy function is predeclared and works for any slice type.
	copy(newProg, codes)
	return newProg
}

func (codes Program) get(address int) (int, Program) {
	if address >= len(codes) {
		//fmt.Println("Reallocating from", len(codes), "to", address*2)
		codes = reallocate(codes, address*2) // Allocated extra
	}
	return codes[address], codes
}

func (codes Program) set(address int, value int) Program {
	if address >= len(codes) {
		//fmt.Println("Reallocating from", len(codes), "to", address*2)
		codes = reallocate(codes, address*2) // Allocated extra
	}
	codes[address] = value
	return codes
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

func getArg(codes Program, ptr int, mode int, base int) (int, Program) {
	var arg int
	arg, codes = codes.get(ptr)
	if mode == 0 {
		arg, codes = codes.get(arg)
	} else if mode == 2 {
		arg, codes = codes.get(base + arg)
	}
	//fmt.Println("getting Arg", ptr, "mode", mode, "base", base, "found", arg)
	return arg, codes
}

func getAddress(codes Program, ptr int, mode int, base int) (int, Program) {
	var arg int
	arg, codes = codes.get(ptr)
	if mode == 2 {
		arg = base + arg
	}
	//fmt.Println("getting Address", ptr, "mode", mode, "base", base, "found", arg)
	return arg, codes
}

const ADD int = 1
const MULT int = 2
const INPUT int = 3
const OUTPUT int = 4
const JUMP_IF_TRUE int = 5
const JUMP_IF_FALSE int = 6
const LESS_THAN int = 7
const EQUALS int = 8
const RELADJUST int = 9
const HALT int = 99

func computer(codes Program, input chan int, output chan int) {
	var (
		//inputPtr int
		ptr      int
		fullcode int
		opcode   int
		relbase  int
		modes    [3]int
	)
	//output = make([]int, 0)
	//fmt.Printf("init addr %p\n", &codes)
	for ptr < len(codes) {
		//fmt.Printf("addr %p\n", &codes)
		fullcode, codes = codes.get(ptr)
		opcode, modes = getOpCode(fullcode)
		switch opcode {
		case ADD:
			var arg1, arg2, out int
			arg1, codes = getArg(codes, ptr+1, modes[0], relbase)
			arg2, codes = getArg(codes, ptr+2, modes[1], relbase)
			out, codes = getAddress(codes, ptr+3, modes[2], relbase)
			codes = codes.set(out, arg1+arg2)
			ptr += 4
		case MULT:
			var arg1, arg2, out int
			arg1, codes = getArg(codes, ptr+1, modes[0], relbase)
			arg2, codes = getArg(codes, ptr+2, modes[1], relbase)
			out, codes = getAddress(codes, ptr+3, modes[2], relbase)
			codes = codes.set(out, arg1*arg2)
			ptr += 4
		case INPUT:
			var arg int
			arg, codes = getAddress(codes, ptr+1, modes[0], relbase)
			val := <-input
			codes = codes.set(arg, val)
			//inputPtr++
			ptr += 2
		case OUTPUT:
			var arg int
			arg, codes = getArg(codes, ptr+1, modes[0], relbase)
			//fmt.Println("Output", arg)
			//output = append(output, arg)
			output <- arg
			ptr += 2
		case JUMP_IF_TRUE:
			var arg1, arg2 int
			arg1, codes = getArg(codes, ptr+1, modes[0], relbase)
			arg2, codes = getArg(codes, ptr+2, modes[1], relbase)
			if arg1 != 0 {
				ptr = arg2
			} else {
				ptr += 3
			}
		case JUMP_IF_FALSE:
			var arg1, arg2 int
			arg1, codes = getArg(codes, ptr+1, modes[0], relbase)
			arg2, codes = getArg(codes, ptr+2, modes[1], relbase)
			if arg1 == 0 {
				ptr = arg2
			} else {
				ptr += 3
			}
		case LESS_THAN:
			var arg1, arg2, out int
			arg1, codes = getArg(codes, ptr+1, modes[0], relbase)
			arg2, codes = getArg(codes, ptr+2, modes[1], relbase)
			out, codes = getAddress(codes, ptr+3, modes[2], relbase)
			if arg1 < arg2 {
				codes = codes.set(out, 1)
			} else {
				codes = codes.set(out, 0)
			}
			ptr += 4
		case EQUALS:
			var arg1, arg2, out int
			arg1, codes = getArg(codes, ptr+1, modes[0], relbase)
			arg2, codes = getArg(codes, ptr+2, modes[1], relbase)
			out, codes = getAddress(codes, ptr+3, modes[2], relbase)
			if arg1 == arg2 {
				codes = codes.set(out, 1)
			} else {
				codes = codes.set(out, 0)
			}
			ptr += 4
		case RELADJUST:
			var arg int
			arg, codes = getArg(codes, ptr+1, modes[0], relbase)
			//fmt.Print("RELADJUST mode ", modes[0], " prev ", relbase, " delta ", arg)
			relbase += arg
			//fmt.Println(" next", relbase)
			ptr += 2
		case HALT:
			//fmt.Println("halt")
			close(output)
			return
		}
	}
	return
}

const up = 0
const right = 1
const down = 2
const left = 3

func move(current [2]int, direction int, turn int) (next [2]int, newdir int) {
	if turn == 0 {
		newdir = (direction - 1 + 4) % 4
	} else {
		newdir = (direction + 1 + 4) % 4
	}
	switch newdir {
	case up:
		next = [2]int{current[0], current[1] + 1}
	case right:
		next = [2]int{current[0] + 1, current[1]}
	case down:
		next = [2]int{current[0], current[1] - 1}
	case left:
		next = [2]int{current[0] - 1, current[1]}
	default:
		next = current
	}
	return
}

func robot(input chan int, output chan int, done chan bool) {
	var (
		direction int
		position  [2]int
		painted   map[[2]int]int
	)
	painted = make(map[[2]int]int)
	for {
		output <- painted[position]
		paint, ok := <-input
		if !ok {
			done <- true
			close(done)
			fmt.Println("Painted", len(painted))
			return
		}
		turn := <-input
		painted[position] = paint
		position, direction = move(position, direction, turn)
	}
}

// Run is the entry point for this solution.
func Run() {
	fmt.Println("Part One")
	var (
		input string
		codes Program
		rtob  chan int
		btor  chan int
		done  chan bool
	)
	rtob = make(chan int, 1)
	btor = make(chan int)
	done = make(chan bool)
	fmt.Scanln(&input)
	codes = toCodes(strings.Split(input, ","))
	go computer(codes, rtob, btor)
	go robot(btor, rtob, done)
	for x := range done {
		fmt.Println("Done?", x)
	}
}
