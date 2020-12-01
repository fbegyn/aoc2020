package main

import (
	"testing"
	"github.com/fbegyn/aoc2020/go/helpers"
)

func BenchmarkPart1(b *testing.B) {
	input := helpers.OpenFile("../../../inputs/day01/input.txt")
	defer input.Close()

	expenses, err  := helpers.LinesToInts(input)
	if err != nil {
		b.Fatal(err)
	}
	
	for i := 0; i < b.N; i++ {
		part1(expenses)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := helpers.OpenFile("../../../inputs/day01/input.txt")
	defer input.Close()

	expenses, err  := helpers.LinesToInts(input)
	if err != nil {
		b.Fatal(err)
	}
	
	for i := 0; i < b.N; i++ {
		part2(expenses)
	}
}

func BenchmarkPart1Robbe(b *testing.B) {
	input := helpers.OpenFile("../../../inputs/day01/robbe.txt")
	defer input.Close()

	expenses, err  := helpers.LinesToInts(input)
	if err != nil {
		b.Fatal(err)
	}
	
	for i := 0; i < b.N; i++ {
		part1(expenses)
	}
}

func BenchmarkPart2Robbe(b *testing.B) {
	input := helpers.OpenFile("../../../inputs/day01/robbe.txt")
	defer input.Close()

	expenses, err  := helpers.LinesToInts(input)
	if err != nil {
		b.Fatal(err)
	}
	
	for i := 0; i < b.N; i++ {
		part2(expenses)
	}
}
