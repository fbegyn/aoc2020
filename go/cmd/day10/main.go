package main

import (
	"log"
	"os"

	"sort"
	"strconv"

	"github.com/fbegyn/aoc2020/go/helpers"
)

func main() {
	file := os.Args[1]
	input := make(chan string, 5)
	go helpers.StreamLines(file, input)

	inputs := []int{0}
	for inp := range input {
		in, err := strconv.Atoi(inp)
		if err != nil {
			log.Fatal(err)
		}
		inputs = append(inputs, in)
	}
	sort.Ints(inputs)
	device := inputs[len(inputs)-1] + 3
	inputs = append(inputs, device)
	differences := make(map[int]int)

	joltage := 0
	for _, rating := range inputs {
		if joltage < rating && rating <= joltage+3 {
			differences[rating-joltage]++
			joltage = rating
		}
	}

	memo := map[int]int{0: 1}

	for i := 1; i < len(inputs); i++ {
		for j := 0; j < i; j++ {
			if inputs[i] - inputs[j] <= 3 {
				memo[i] += memo[j]
			}
		}
	}

	log.Printf("solution for part 1: %d", differences[1]*differences[3])
	log.Printf("solution for part 2: %v", memo)
}
