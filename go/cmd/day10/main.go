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

	inputs := []int{}
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
			differences[rating-joltage] += 1
			joltage = rating
		}
	}

	

	for diff, freq := range differences {
		log.Printf("There are %d differences of %d joltage", freq, diff)
	}

	log.Printf("solution for part 1: %d", differences[1]*differences[3])
	log.Printf("solution for part 2: %v", Arrange(inputs, 0))
}

func Arrange(path []int, ind int) int {
	if ind >= len(path) {
		return 1
	}

	x := path[ind]
	ans := Arrange(path, ind+1)

	if ind+2 < len(path) && path[ind+2]-x <= 3 {
		ans += Arrange(path, ind+2)
		if ind+3 < len(path) && path[ind+3]-x <= 3 {
			ans += Arrange(path, ind+3)
		}
	}
	return ans
}

