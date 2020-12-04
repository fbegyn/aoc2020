package main

import (
	"log"
	"os"

	"github.com/fbegyn/aoc2020/go/helpers"
	"errors"
)

func main() {
	file := os.Args[1]
	input := helpers.OpenFile(file)
	defer input.Close()

	expenses, err  := helpers.LinesToInts(input)
	if err != nil {
		log.Fatal(err)
	}
	
	part1, err := part1(expenses)
	if err != nil {
		log.Fatal(err)
	}
	part2, err := part2(expenses)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("answer for part 1 is %v", part1)
	log.Printf("answer for part 2 is %v", part2)
}

func part1(expenses []int) (int, error) {
	for i, v := range expenses {
		for _, w := range expenses[i+1:] {
			if v+w == 2020 {
				return v * w, nil
			}
		}
	}
	return 0, errors.New("no solution found for part 2")
}

func part2(expenses []int) (int, error) {
	t := helpers.MinInt(expenses)
	for i, v := range expenses {
		for j, w := range expenses[i+1:] {
			if v+w > 2020 - t {
				continue
			}
			for _, z := range expenses[j+1:] {
				if v+w+z == 2020 {
					return v * w * z, nil
				}
			}
		}
	}
	return 0, errors.New("no solution found for part 2")
}
