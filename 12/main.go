package main

import (
	"flag"
	"fmt"
	"os"

	"./one"
	"./two"
)

func main() {
	var (
		a bool
		b bool
	)
	flag.BoolVar(&a, "one", false, "Execute part one.")
	flag.BoolVar(&b, "two", false, "Execute part two.")
	flag.Parse()
	if a {
		one.Run()
	} else if b {
		two.Run()
	} else {
		fmt.Println("Usage:")
		fmt.Println(os.Args[0], "<option>")
		flag.PrintDefaults()
	}
}
