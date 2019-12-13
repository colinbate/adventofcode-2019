package computer

import (
	"fmt"
	"strconv"
	"strings"
)

// Program is a sequence of instructinos for intcode computer
type Program []int

func reallocate(codes Program, newSize int) Program {
	newProg := make(Program, newSize)
	copy(newProg, codes)
	return newProg
}

func (codes *Program) get(address int) int {
	if address >= len(*codes) {
		//fmt.Println("Reallocating from", len(codes), "to", address*2)
		*codes = reallocate(*codes, address*2) // Allocated extra
	}
	return (*codes)[address]
}

func (codes *Program) set(address int, value int) {
	if address >= len(*codes) {
		//fmt.Println("Reallocating from", len(codes), "to", address*2)
		*codes = reallocate(*codes, address*2) // Allocated extra
	}
	(*codes)[address] = value
}

func getOpCode(code int) (int, [3]int) {
	var modes [3]int
	op := code % 100
	modes[0] = code / 100 % 10
	modes[1] = code / 1000 % 10
	modes[2] = code / 10000 % 10
	return op, modes
}

func (codes *Program) getArg(ptr int, mode int, base int) int {
	var arg int
	arg = codes.get(ptr)
	if mode == 0 {
		arg = codes.get(arg)
	} else if mode == 2 {
		arg = codes.get(base + arg)
	}
	//fmt.Println("getting Arg", ptr, "mode", mode, "base", base, "found", arg)
	return arg
}

func (codes *Program) getAddress(ptr int, mode int, base int) int {
	var arg int
	arg = codes.get(ptr)
	if mode == 2 {
		arg = base + arg
	}
	//fmt.Println("getting Address", ptr, "mode", mode, "base", base, "found", arg)
	return arg
}

const opADD int = 1
const opMULT int = 2
const opINPUT int = 3
const opOUTPUT int = 4
const opJUMPIFTRUE int = 5
const opJUMPIFFALSE int = 6
const opLESSTHAN int = 7
const opEQUALS int = 8
const opRELADJUST int = 9
const opHALT int = 99

// Execute will run the program provided
func (codes *Program) Execute(input chan int, output chan int) {
	var (
		ptr     int
		opcode  int
		relbase int
		modes   [3]int
	)
	for ptr < len(*codes) {
		opcode, modes = getOpCode(codes.get(ptr))
		switch opcode {
		case opADD:
			var arg1, arg2, out int
			arg1 = codes.getArg(ptr+1, modes[0], relbase)
			arg2 = codes.getArg(ptr+2, modes[1], relbase)
			out = codes.getAddress(ptr+3, modes[2], relbase)
			codes.set(out, arg1+arg2)
			ptr += 4
		case opMULT:
			var arg1, arg2, out int
			arg1 = codes.getArg(ptr+1, modes[0], relbase)
			arg2 = codes.getArg(ptr+2, modes[1], relbase)
			out = codes.getAddress(ptr+3, modes[2], relbase)
			codes.set(out, arg1*arg2)
			ptr += 4
		case opINPUT:
			var arg int
			arg = codes.getAddress(ptr+1, modes[0], relbase)
			val := <-input
			codes.set(arg, val)
			ptr += 2
		case opOUTPUT:
			var arg int
			arg = codes.getArg(ptr+1, modes[0], relbase)
			output <- arg
			ptr += 2
		case opJUMPIFTRUE:
			var arg1, arg2 int
			arg1 = codes.getArg(ptr+1, modes[0], relbase)
			arg2 = codes.getArg(ptr+2, modes[1], relbase)
			if arg1 != 0 {
				ptr = arg2
			} else {
				ptr += 3
			}
		case opJUMPIFFALSE:
			var arg1, arg2 int
			arg1 = codes.getArg(ptr+1, modes[0], relbase)
			arg2 = codes.getArg(ptr+2, modes[1], relbase)
			if arg1 == 0 {
				ptr = arg2
			} else {
				ptr += 3
			}
		case opLESSTHAN:
			var arg1, arg2, out int
			arg1 = codes.getArg(ptr+1, modes[0], relbase)
			arg2 = codes.getArg(ptr+2, modes[1], relbase)
			out = codes.getAddress(ptr+3, modes[2], relbase)
			if arg1 < arg2 {
				codes.set(out, 1)
			} else {
				codes.set(out, 0)
			}
			ptr += 4
		case opEQUALS:
			var arg1, arg2, out int
			arg1 = codes.getArg(ptr+1, modes[0], relbase)
			arg2 = codes.getArg(ptr+2, modes[1], relbase)
			out = codes.getAddress(ptr+3, modes[2], relbase)
			if arg1 == arg2 {
				codes.set(out, 1)
			} else {
				codes.set(out, 0)
			}
			ptr += 4
		case opRELADJUST:
			var arg int
			arg = codes.getArg(ptr+1, modes[0], relbase)
			relbase += arg
			ptr += 2
		case opHALT:
			//fmt.Println("halt")
			close(output)
			return
		}
	}
	return
}

func toCodes(vals []string) Program {
	codes := make([]int, len(vals))
	for i, s := range vals {
		codes[i], _ = strconv.Atoi(s)
	}
	return codes
}

// LoadProgram will load a program from stdin
func LoadProgram() Program {
	var input string
	fmt.Scanln(&input)
	codes := toCodes(strings.Split(input, ","))
	return codes
}
