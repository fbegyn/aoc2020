package main

import (
	"bufio"
	"testing"

	"github.com/fbegyn/aoc2020/go/helpers"
)

func BenchmarkPart1(b *testing.B) {
	input := helpers.OpenFile("../../../inputs/day03/input.txt")
	defer input.Close()

	routeMap := [][]rune{}
	rowCount := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, len(line))
		for ind, ch := range line {
			row[ind] = ch
		}
		routeMap = append(routeMap, row)
		rowCount += 1
	}

	position := NewLocation(routeMap)
	for i := 0; i < b.N; i++ {
	    part1(position)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := helpers.OpenFile("../../../inputs/day03/input.txt")
	defer input.Close()

	routeMap := [][]rune{}
	rowCount := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, len(line))
		for ind, ch := range line {
			row[ind] = ch
		}
		routeMap = append(routeMap, row)
		rowCount += 1
	}

	position := NewLocation(routeMap)
	for i := 0; i < b.N; i++ {
	    part2(position)
	}
}
