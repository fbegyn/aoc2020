package main

import (
	"os"
	"strconv"
	"log"
	"strings"
	"sort"

	"github.com/fbegyn/aoc2020/go/helpers"
)

func main() {
	file := os.Args[1]
	input := make(chan string, 5)
	go helpers.StreamLines(file, input)

	ids := []int{}

	for line := range input {
		line = strings.ReplaceAll(line, "F", "0")
		line = strings.ReplaceAll(line, "B", "1")
		line = strings.ReplaceAll(line, "L", "0")
		line = strings.ReplaceAll(line, "R", "1")
		ID, err := strconv.ParseInt(line, 2, 32)
		if err != nil {
			log.Fatalf("failed to parse ID: %v", err)
		}
		ids = append(ids, int(ID))
	}

	// part 1
	sort.Ints(ids)

	// part 2
	var mySeat int
	for ind := range ids {
		if mySeat =ids[ind]+1; mySeat != ids[ind+1] {
			break
		}
	}

	log.Printf("solution for part 1: %d\n", ids[len(ids)-1])
	log.Printf("solution for part 2: %d\n", mySeat)
}

func BinaryParse(input string, lower, upper int, output chan<- int) {
	for _, r := range input {
		switch r {
		case 'F':
			upper = lower + (upper - lower) / 2
		case 'B':
			lower = lower + (upper -lower) / 2 +1
		case 'L':
			upper = lower + (upper - lower) / 2
		case 'R':
			lower = lower + (upper -lower) / 2 +1
		}
	}
	if upper == lower {
		output <- upper
	}
}
