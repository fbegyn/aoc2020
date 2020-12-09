package main

import (
	"log"
	"os"
	"strconv"
	"sort"

	"github.com/fbegyn/aoc2020/go/helpers"
)

func main() {
	file := os.Args[1]
	input := make(chan string, 5)
	go helpers.StreamLines(file, input)

	in := []int{}
	preambleSize, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	for inp := range input {
		number, err := strconv.Atoi(inp)
		if err != nil {
			log.Fatal(err)
		}
		in = append(in, number)
	}

	invalid := 0
	index := 0
	for invalid == 0 && index+preambleSize < len(in) {
		preamble := make([]int, preambleSize)
		copy(preamble, in[index:index+preambleSize])
		sort.Ints(preamble)
		min := preamble[0]+preamble[1]
		max := preamble[preambleSize-1]+preamble[preambleSize-2]
		nr := in[index+preambleSize]
		if nr < min || max < nr {
			invalid = nr
			continue
		}
		sum := false
		for i, a := range preamble {
			for _, b := range preamble[i:] {
				if a + b == nr {
					sum = true
					break
				}
			}
		}
		if !sum {
			invalid = nr
			continue
		}
		index += 1
	}

	indexLow, indexHigh := 0, 0
	found := false
	for i := 2; !found; i++ {
		for j := range in {
			if invalid == helpers.SumOfIntArray(in[j:j+i]) {
				indexLow, indexHigh = j, j+i
				found = true
				break
			}
		}
	}

	list := make([]int, indexHigh-indexLow)
	copy(list, in[indexLow:indexHigh])
	sort.Ints(list)

	log.Printf("solution to part 1: %d", invalid)
	log.Printf("solution to part 2: %d", list[0]+list[len(list)-1])
}
