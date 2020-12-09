package main

import (
	"sort"
	"strconv"
	"testing"

	"github.com/fbegyn/aoc2020/go/helpers"
)

func BenchmarkPart1(b *testing.B) {
	file := "../../../inputs/day09/input.txt"
	input := make(chan string, 5)
	go helpers.StreamLines(file, input)

	in := []int{}
	for inp := range input {
		number, err := strconv.Atoi(inp)
		if err != nil {
			b.Fatal(err)
		}
		in = append(in, number)
	}

	for i := 0; i < b.N; i++ {
		part1(in, 25)
	}
}

func BenchmarkPart2(b *testing.B) {
	file := "../../../inputs/day09/input.txt"
	input := make(chan string, 5)
	go helpers.StreamLines(file, input)

	in := []int{}
	for inp := range input {
		number, err := strconv.Atoi(inp)
		if err != nil {
			b.Fatal(err)
		}
		in = append(in, number)
	}

	for i := 0; i < b.N; i++ {
		indexLow, indexHigh := part2(in, 1492208709)
		sort.Ints(in[indexLow:indexHigh-1])
	}
}
