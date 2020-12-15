package main

import (
	"log"
	"os"
	"strings"

	"github.com/fbegyn/aoc2020/go/helpers"
)

func main() {
	file := os.Args[1]
	input := make(chan string, 5)
	go helpers.StreamLines(file, input)

	changes := []int{}
	index := 0

	instructions := []string{}
	for inp := range input {
		instructions = append(instructions, inp)
		if opcode := strings.Split(inp, " ")[0]; opcode == "jmp" || opcode == "nop" {
			changes = append(changes, index)
		}
		index += 1
	}

	output := make(chan int)
	halt := make(chan bool)
	loop := make(chan bool)

	go helpers.RunProgram(instructions, output, halt, loop)

	select {
	case <-loop:
		log.Printf("solution to part 1: %d", <-output)
	case <-halt:
		log.Printf("solution to part 2: %d", <-output)
	}

	for _, ch := range changes {
		fix := helpers.ToggleInstruction(instructions, ch)
		go helpers.RunProgram(fix, output, halt, loop)
	}


	<-halt
	log.Printf("solution to part 2: %d", <-output)
}
