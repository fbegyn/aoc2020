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

	freq := make(map[rune]int)
	people := 0
	counts1 := []int{}
	counts2 := []int{}

	for inp := range input {
		inp = strings.TrimSpace(inp)

		// reset the freuency map on empty lines (new groups)
		if inp == "" {
			counts1 = append(counts1, SumMap(freq))
			counts2 = append(counts2, CheckMap(freq, people))

			people = 0
			freq = make(map[rune]int)
			continue
		}

		// compose frequency map for the current line
		for _, r := range inp {
			freq[r] += 1
		}
		people += 1
	}
	counts1 = append(counts1, SumMap(freq))
	counts2 = append(counts2, CheckMap(freq, people))

	sum1 := 0
	sum2 := 0
	for ind := range counts1 {
		sum1 += counts1[ind]
		sum2 += counts2[ind]
	}

	log.Printf("solution to part 1: %d\n", sum1)
	log.Printf("solution to part 2: %d\n", sum2)
}

func CheckMap(m map[rune]int, n int) (sum int) {
	for _, freq := range m {
		if freq == n {
			sum += 1
		}
	}
	return
}

func SumMap(m map[rune]int) (sum int) {
	for range m {
		sum += 1
	}
	return
}
