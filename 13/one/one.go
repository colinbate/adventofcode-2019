package one

import (
	"fmt"

	"../computer"
)

const empty = 0
const wall = 1
const block = 2
const paddle = 3
const ball = 4

func arcade(instructions chan int, done chan bool) {
	tiles := make(map[[2]int]int)
	count := 0
	for {
		x, ok := <-instructions
		if !ok {
			for _, t := range tiles {
				if t == block {
					count++
				}
			}
			fmt.Println("Blocks:", count)
			close(done)
			return
		}
		y := <-instructions
		tid := <-instructions
		tiles[[2]int{x, y}] = tid
	}
}

// Run is the entry point for this solution.
func Run() {
	fmt.Println("Part One")
	game := computer.LoadProgram()
	io := make(chan int)
	done := make(chan bool)
	go game.Execute(make(chan int), io)
	go arcade(io, done)
	<-done
}
